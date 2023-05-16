package pfcp

import (
	"upmf/internal/pfcp/pfcpmsg"
	"upmf/internal/pfcp/pfcptypes"
	"net"
)

type AssociationProducer interface {
	HandleAssociationSetupRequest(string, *pfcpmsg.PFCPAssociationSetupRequest) (*pfcpmsg.PFCPAssociationSetupResponse, error)
	HandleAssociationReleaseRequest(string, *pfcpmsg.PFCPAssociationReleaseRequest) (*pfcpmsg.PFCPAssociationReleaseResponse, error)
	HandleHeartbeatRequest(string, *pfcpmsg.HeartbeatRequest) (*pfcpmsg.HeartbeatResponse, error)
}
type SessionProducer interface {
	HandleSessionReportRequest(string, uint64, *pfcpmsg.PFCPSessionReportRequest) (*pfcpmsg.PFCPSessionReportResponse, uint64, error)
}

type Producer interface {
	AssociationProducer
	SessionProducer
}

func (proto *Pfcp) handleReq(remote *net.UDPAddr, msg *pfcpmsg.Message) {
	log.Debugf("receive a request of type %d from %s", msg.Header.MessageType, remote)
	if infoinf := proto.queue.find(remote.String(), msg.Header.SequenceNumber); infoinf != nil {
		//duplicated request
		if info, ok := infoinf.(*RspSendingInfo); ok {
			//resend the response
			log.Warnf("re-send the response to %s", info.remote.String())
			proto.fwd.WriteTo(info.msg, remote)
		}
	}
	//a new request
	switch msg.Header.MessageType {

	case pfcpmsg.PFCP_HEARTBEAT_REQUEST:
		proto.handleHeartbeatReq(remote, msg)

	case pfcpmsg.PFCP_PFD_MANAGEMENT_REQUEST:

	case pfcpmsg.PFCP_ASSOCIATION_SETUP_REQUEST:
		proto.handleAssSetReq(remote, msg)

	case pfcpmsg.PFCP_ASSOCIATION_UPDATE_REQUEST:

	case pfcpmsg.PFCP_ASSOCIATION_RELEASE_REQUEST:
		proto.handleAssRelReq(remote, msg)

	case pfcpmsg.PFCP_NODE_REPORT_REQUEST:
		//proto.handleNodeRepReq(remote, msg)

	case pfcpmsg.PFCP_SESSION_SET_DELETION_REQUEST:

	case pfcpmsg.PFCP_SESSION_ESTABLISHMENT_REQUEST:

	case pfcpmsg.PFCP_SESSION_MODIFICATION_REQUEST:

	case pfcpmsg.PFCP_SESSION_DELETION_REQUEST:

	case pfcpmsg.PFCP_SESSION_REPORT_REQUEST:
		proto.handleSessRepReq(remote, msg)

	default:
	}
}

func (proto *Pfcp) handleAssSetReq(remote *net.UDPAddr, msg *pfcpmsg.Message) {
	req := msg.Body.(pfcpmsg.PFCPAssociationSetupRequest)
	if body, err := proto.ahandler.HandleAssociationSetupRequest(remote.String(), &req); err == nil {
		body.RecoveryTimeStamp = &pfcptypes.RecoveryTimeStamp{
			RecoveryTimeStamp: proto.fwd.When(),
		}
		body.NodeID = proto.ctx.NodeId()

		rsp := &pfcpmsg.Message{
			Header: pfcpmsg.Header{
				Version:        pfcpmsg.PfcpVersion,
				MP:             0,
				S:              pfcpmsg.SEID_NOT_PRESENT,
				MessageType:    pfcpmsg.PFCP_ASSOCIATION_SETUP_RESPONSE,
				SequenceNumber: msg.Header.SequenceNumber,
			},
			Body: body,
		}
		proto.sendRsp(rsp, remote)
	}
}

func (proto *Pfcp) handleAssRelReq(remote *net.UDPAddr, msg *pfcpmsg.Message) {
	req := msg.Body.(pfcpmsg.PFCPAssociationReleaseRequest)
	if body, err := proto.ahandler.HandleAssociationReleaseRequest(remote.String(), &req); err == nil {
		body.NodeID = proto.ctx.NodeId()
		rsp := &pfcpmsg.Message{
			Header: pfcpmsg.Header{
				Version:        pfcpmsg.PfcpVersion,
				MP:             0,
				S:              pfcpmsg.SEID_NOT_PRESENT,
				MessageType:    pfcpmsg.PFCP_ASSOCIATION_RELEASE_RESPONSE,
				SequenceNumber: msg.Header.SequenceNumber,
			},
			Body: body,
		}
		proto.sendRsp(rsp, remote)
	}
}

func (proto *Pfcp) handleSessRepReq(remote *net.UDPAddr, msg *pfcpmsg.Message) {
	req := msg.Body.(pfcpmsg.PFCPSessionReportRequest)
	if body, seid, err := proto.shandler.HandleSessionReportRequest(remote.String(), msg.Header.SEID, &req); err == nil {
		rsp := &pfcpmsg.Message{
			Header: pfcpmsg.Header{
				Version:        pfcpmsg.PfcpVersion,
				MP:             0,
				S:              pfcpmsg.SEID_PRESENT,
				MessageType:    pfcpmsg.PFCP_SESSION_REPORT_RESPONSE,
				SequenceNumber: msg.Header.SequenceNumber,
				SEID:           seid,
			},
			Body: body,
		}
		proto.sendRsp(rsp, remote)
	}

}

func (proto *Pfcp) handleHeartbeatReq(remote *net.UDPAddr, msg *pfcpmsg.Message) {
	req := msg.Body.(pfcpmsg.HeartbeatRequest)
	if body, err := proto.ahandler.HandleHeartbeatRequest(remote.String(), &req); err == nil {
		body.RecoveryTimeStamp = &pfcptypes.RecoveryTimeStamp{
			RecoveryTimeStamp: proto.fwd.When(),
		}

		rsp := &pfcpmsg.Message{
			Header: pfcpmsg.Header{
				Version:        pfcpmsg.PfcpVersion,
				MP:             0,
				S:              pfcpmsg.SEID_NOT_PRESENT,
				MessageType:    pfcpmsg.PFCP_HEARTBEAT_RESPONSE,
				SequenceNumber: msg.Header.SequenceNumber,
			},
			Body: body,
		}
		proto.sendRsp(rsp, remote)
	}
}
