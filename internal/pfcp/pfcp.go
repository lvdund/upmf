package pfcp

import (
	"upmf/internal/pfcp/pfcpmsg"
	"upmf/internal/pfcp/pfcptypes"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

const (
	PFCP_MSG_RECV_QUEUE_SIZE int = 1024 //size of the queue to receive Pfcp message from forwwarder
	PFCP_MSG_SEND_QUEUE_SIZE int = 1024 //size of the queue to receive Pfcp message for sending to forwwarder
)

type Endpoint interface {
	UdpAddr() *net.UDPAddr
}

type PfcpSession interface {
	Endpoint
	RemoteSeid() uint64
	LocalSeid() uint64
	FillDeletionRequest(*pfcpmsg.PFCPSessionDeletionRequest)
	FillEstablishmentRequest(*pfcpmsg.PFCPSessionEstablishmentRequest)
	FillModificationRequest(*pfcpmsg.PFCPSessionModificationRequest)
}

type NodeContext interface {
	NodeId() *pfcptypes.NodeID
	Port() int
	SessionHandler() SessionProducer
	AssociationHandler() AssociationProducer
}

type Pfcp struct {
	ctx      NodeContext
	fwd      *Forwarder
	seq      uint32 //for generating request sequence number
	wg       sync.WaitGroup
	shandler SessionProducer     //upper handler to response messages
	ahandler AssociationProducer //upper handler to response messages
	done     chan bool           //trigger termination in child loops
	sending  chan *ReqSendingInfo
	queue    ExpiringList
}

func NewPfcp(ctx NodeContext) *Pfcp {
	id := ctx.NodeId()
	addr := net.UDPAddr{
		IP:   id.ResolveNodeIdToIp(),
		Port: ctx.Port(),
	}
	ret := &Pfcp{
		ctx:      ctx,
		fwd:      newForwarder(addr),
		sending:  make(chan *ReqSendingInfo, PFCP_MSG_SEND_QUEUE_SIZE),
		done:     make(chan bool),
		queue:    newExpiringList(),
		shandler: ctx.SessionHandler(),
		ahandler: ctx.AssociationHandler(),
	}
	return ret
}

func (proto *Pfcp) Start() (err error) {
	recv := make(chan RecvInfo, PFCP_MSG_RECV_QUEUE_SIZE)
	if err = proto.fwd.start(recv); err == nil {
		go proto.receivingloop(recv)
		go proto.sendingloop()
	}
	return
}

func (proto *Pfcp) Stop() {
	proto.fwd.stop()
	close(proto.done)
	proto.wg.Wait()
}

// waits for request messages to send
func (proto *Pfcp) sendingloop() {
	proto.wg.Add(1)
	defer proto.wg.Done()
	ticker := time.NewTicker(EXPIRING_CHECK_INTERVAL)
	for {
		select {
		case <-proto.done:
			return
		case <-ticker.C:
			proto.queue.flush()

		case info := <-proto.sending:
			//sending the message
			if info.err = proto.fwd.WriteTo(info.msg, info.remote); info.err != nil {
				//terminate sending
				close(info.done)
			} else {
				//cache the request for resending (in case of timeout)  and
				//searching (in case of receiving a response)
				proto.queue.add(info)
			}
		}
	}
}

// receiving messages from forwarder
func (proto *Pfcp) receivingloop(recv chan RecvInfo) {
	proto.wg.Add(1)
	defer proto.wg.Done()
	for info := range recv {
		proto.handle(info.remote, info.msg)
	}

}

func (proto *Pfcp) handle(remote *net.UDPAddr, msg *pfcpmsg.Message) {
	if msg.IsRequest() {
		proto.handleReq(remote, msg)
	} else {
		proto.handleRsp(remote, msg)
	}
}

func (proto *Pfcp) handleRsp(remote *net.UDPAddr, msg *pfcpmsg.Message) {
	log.Debugf("receive a response of type %d from %s", msg.Header.MessageType, remote)
	if infoinf := proto.queue.pop(remote.String(), msg.Header.SequenceNumber); infoinf != nil {
		if info, ok := infoinf.(*ReqSendingInfo); ok {
			match := false
			switch msg.Header.MessageType {
			case pfcpmsg.PFCP_HEARTBEAT_RESPONSE:
				match = info.msg.Header.MessageType == pfcpmsg.PFCP_HEARTBEAT_REQUEST
			case pfcpmsg.PFCP_PFD_MANAGEMENT_RESPONSE:
				match = info.msg.Header.MessageType == pfcpmsg.PFCP_PFD_MANAGEMENT_REQUEST
			case pfcpmsg.PFCP_ASSOCIATION_SETUP_RESPONSE:
				match = info.msg.Header.MessageType == pfcpmsg.PFCP_ASSOCIATION_SETUP_REQUEST
			case pfcpmsg.PFCP_ASSOCIATION_UPDATE_RESPONSE:
				match = info.msg.Header.MessageType == pfcpmsg.PFCP_ASSOCIATION_UPDATE_REQUEST
			case pfcpmsg.PFCP_ASSOCIATION_RELEASE_RESPONSE:
				match = info.msg.Header.MessageType == pfcpmsg.PFCP_ASSOCIATION_RELEASE_REQUEST
			case pfcpmsg.PFCP_NODE_REPORT_RESPONSE:
				match = info.msg.Header.MessageType == pfcpmsg.PFCP_NODE_REPORT_REQUEST
			case pfcpmsg.PFCP_SESSION_SET_DELETION_RESPONSE:
				match = info.msg.Header.MessageType == pfcpmsg.PFCP_SESSION_SET_DELETION_REQUEST
			case pfcpmsg.PFCP_SESSION_ESTABLISHMENT_RESPONSE:
				match = info.msg.Header.MessageType == pfcpmsg.PFCP_SESSION_ESTABLISHMENT_REQUEST
			case pfcpmsg.PFCP_SESSION_MODIFICATION_RESPONSE:
				match = info.msg.Header.MessageType == pfcpmsg.PFCP_SESSION_MODIFICATION_REQUEST
			case pfcpmsg.PFCP_SESSION_DELETION_RESPONSE:
				match = info.msg.Header.MessageType == pfcpmsg.PFCP_SESSION_DELETION_REQUEST
			case pfcpmsg.PFCP_SESSION_REPORT_RESPONSE:
				match = info.msg.Header.MessageType == pfcpmsg.PFCP_SESSION_REPORT_REQUEST
			default:
			}

			if !match {
				info.err = fmt.Errorf("Mismatched response")
			} else {
				info.rsp = msg
			}
			//terminate the sending task
			close(info.done)
		}
	}
}

// send a request then wait for a response.
func (proto *Pfcp) sendReq(msg *pfcpmsg.Message, remote *net.UDPAddr) (rsp *pfcpmsg.Message, err error) {
	info := newReqSendingInfo(msg, remote, proto.scheduleReqSending)
	//schedule for sending
	proto.scheduleReqSending(info)
	//wait for a response
	<-info.done
	rsp = info.rsp
	err = info.err
	return
}

func (proto *Pfcp) sendRsp(msg *pfcpmsg.Message, remote *net.UDPAddr) (err error) {
	if err = proto.fwd.WriteTo(msg, remote); err == nil {
		//cache the message for a certain time duration to resend in cases
		//where duplicated requests arrive
		proto.queue.add(newRspSendingInfo(msg, remote))
	}

	return
}

func (proto *Pfcp) scheduleReqSending(info *ReqSendingInfo) {
	//increate sending retry counter for the request
	info.retry++
	//log.Infof("schedule for sending %d", info.retry)
	//push to sending chanel
	proto.sending <- info
}

// generate sequence number for PFCP sending request
func (proto *Pfcp) sequenceNumber() uint32 {
	return atomic.AddUint32(&proto.seq, 1)
}
