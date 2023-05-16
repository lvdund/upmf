package pfcp

import (
	"upmf/internal/pfcp/pfcpmsg"
	"upmf/internal/pfcp/pfcptypes"
	"fmt"
)

func (proto *Pfcp) SendPfcpAssociationSetupRequest(remote Endpoint) (rsp *pfcpmsg.PFCPAssociationSetupResponse, err error) {
	req := &pfcpmsg.PFCPAssociationSetupRequest{
		NodeID: proto.ctx.NodeId(),
		RecoveryTimeStamp: &pfcptypes.RecoveryTimeStamp{
			RecoveryTimeStamp: proto.fwd.When(),
		},
		CPFunctionFeatures: &pfcptypes.CPFunctionFeatures{
			SupportedFeatures: 0,
		},
	}

	reqmsg := &pfcpmsg.Message{
		Header: pfcpmsg.Header{
			Version:        pfcpmsg.PfcpVersion,
			MP:             0,
			S:              pfcpmsg.SEID_NOT_PRESENT,
			MessageType:    pfcpmsg.PFCP_ASSOCIATION_SETUP_REQUEST,
			SequenceNumber: proto.sequenceNumber(),
		},
		Body: req,
	}
	var rspmsg *pfcpmsg.Message
	if rspmsg, err = proto.sendReq(reqmsg, remote.UdpAddr()); err == nil {
		body := rspmsg.Body.(pfcpmsg.PFCPAssociationSetupResponse)
		rsp = &body
	}
	return
}

func (proto *Pfcp) SendPfcpAssociationReleaseRequest(remote Endpoint) (rsp *pfcpmsg.PFCPAssociationReleaseResponse, err error) {
	req := &pfcpmsg.PFCPAssociationReleaseRequest{
		NodeID: proto.ctx.NodeId(),
	}

	reqmsg := &pfcpmsg.Message{
		Header: pfcpmsg.Header{
			Version:        pfcpmsg.PfcpVersion,
			MP:             0,
			S:              pfcpmsg.SEID_NOT_PRESENT,
			MessageType:    pfcpmsg.PFCP_ASSOCIATION_RELEASE_REQUEST,
			SequenceNumber: proto.sequenceNumber(),
		},
		Body: req,
	}

	var rspmsg *pfcpmsg.Message
	if rspmsg, err = proto.sendReq(reqmsg, remote.UdpAddr()); err == nil {
		body := rspmsg.Body.(pfcpmsg.PFCPAssociationReleaseResponse)
		rsp = &body
	}
	return
}

func (proto *Pfcp) SendPfcpSessionDeletionRequest(session PfcpSession) (rsp *pfcpmsg.PFCPSessionDeletionResponse, err error) {
	reqbody := &pfcpmsg.PFCPSessionDeletionRequest{}
	reqmsg := &pfcpmsg.Message{
		Header: pfcpmsg.Header{
			Version:         pfcpmsg.PfcpVersion,
			MP:              1,
			S:               pfcpmsg.SEID_PRESENT,
			MessageType:     pfcpmsg.PFCP_SESSION_DELETION_REQUEST,
			SEID:            session.RemoteSeid(),
			SequenceNumber:  proto.sequenceNumber(),
			MessagePriority: 12,
		},
		Body: reqbody,
	}
	session.FillDeletionRequest(reqbody)
	var rspmsg *pfcpmsg.Message
	if rspmsg, err = proto.sendReq(reqmsg, session.UdpAddr()); err == nil {
		if rspmsg.Header.SEID == session.LocalSeid() {
			body := rspmsg.Body.(pfcpmsg.PFCPSessionDeletionResponse)
			rsp = &body
		} else {
			err = fmt.Errorf("mismatched SEID")
		}
	}
	return
}

func (proto *Pfcp) SendPfcpHeartbeatRequest(remote Endpoint) (rsp *pfcpmsg.HeartbeatResponse, err error) {
	req := &pfcpmsg.HeartbeatRequest{
		RecoveryTimeStamp: &pfcptypes.RecoveryTimeStamp{
			RecoveryTimeStamp: proto.fwd.When(),
		},
	}

	reqmsg := &pfcpmsg.Message{
		Header: pfcpmsg.Header{
			Version:        pfcpmsg.PfcpVersion,
			MP:             0,
			S:              pfcpmsg.SEID_NOT_PRESENT,
			MessageType:    pfcpmsg.PFCP_HEARTBEAT_REQUEST,
			SequenceNumber: proto.sequenceNumber(),
		},
		Body: req,
	}
	var rspmsg *pfcpmsg.Message
	if rspmsg, err = proto.sendReq(reqmsg, remote.UdpAddr()); err == nil {
		body := rspmsg.Body.(pfcpmsg.HeartbeatResponse)
		rsp = &body
	}

	return
}

func (proto *Pfcp) SendPfcpSessionEstablishmentRequest(session PfcpSession) (rsp *pfcpmsg.PFCPSessionEstablishmentResponse, err error) {
	reqbody := &pfcpmsg.PFCPSessionEstablishmentRequest{
		NodeID: proto.ctx.NodeId(),
	}
	reqmsg := &pfcpmsg.Message{
		Header: pfcpmsg.Header{
			Version:         pfcpmsg.PfcpVersion,
			MP:              1,
			S:               pfcpmsg.SEID_PRESENT,
			MessageType:     pfcpmsg.PFCP_SESSION_ESTABLISHMENT_REQUEST,
			SEID:            0, /*session.RemoteSeid()*/
			SequenceNumber:  proto.sequenceNumber(),
			MessagePriority: 0,
		},
		Body: reqbody,
	}

	session.FillEstablishmentRequest(reqbody)

	var rspmsg *pfcpmsg.Message
	if rspmsg, err = proto.sendReq(reqmsg, session.UdpAddr()); err == nil {
		if rspmsg.Header.SEID == session.LocalSeid() {
			body := rspmsg.Body.(pfcpmsg.PFCPSessionEstablishmentResponse)
			rsp = &body
		} else {
			err = fmt.Errorf("mismatched SEID")
		}
	}

	return
}

func (proto *Pfcp) SendPfcpSessionModificationRequest(session PfcpSession) (rsp *pfcpmsg.PFCPSessionModificationResponse, err error) {
	reqbody := &pfcpmsg.PFCPSessionModificationRequest{}
	reqmsg := &pfcpmsg.Message{
		Header: pfcpmsg.Header{
			Version:         pfcpmsg.PfcpVersion,
			MP:              1,
			S:               pfcpmsg.SEID_PRESENT,
			MessageType:     pfcpmsg.PFCP_SESSION_MODIFICATION_REQUEST,
			SEID:            session.RemoteSeid(),
			SequenceNumber:  proto.sequenceNumber(),
			MessagePriority: 12,
		},
		Body: reqbody,
	}

	session.FillModificationRequest(reqbody)

	var rspmsg *pfcpmsg.Message
	if rspmsg, err = proto.sendReq(reqmsg, session.UdpAddr()); err == nil {
		if rspmsg.Header.SEID == session.LocalSeid() {
			body := rspmsg.Body.(pfcpmsg.PFCPSessionModificationResponse)
			rsp = &body
		} else {
			err = fmt.Errorf("mismatched SEID")
		}
	}

	return
}
