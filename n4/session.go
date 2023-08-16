package n4

import "net"

/*
type SmfContext interface {
}
*/
type MsgSession struct {
	upf *Upf //the session is on this upf
	//	smcontext  SmfContext //the session belongs to this SmContext
	localseid  uint64 // session identification in PFCP received messages
	remoteseid uint64 //session identification in PFCP sending messages
	pdrs       []*PDR
}

func newMsgSession(localseid uint64, upf *Upf) *MsgSession {
	return &MsgSession{
		localseid: localseid,
		upf:       upf,
	}
}

func (s *MsgSession) UdpAddr() *net.UDPAddr {
	return &net.UDPAddr{
		IP:   s.upf.ip,
		Port: int(s.upf.port),
	}
}

func (s *MsgSession) LocalSeid() uint64 {
	return s.localseid
}
func (s *MsgSession) SetRemoteSeid(seid uint64) {
	s.remoteseid = seid
}

func (s *MsgSession) RemoteSeid() uint64 {
	return s.remoteseid
}

func (s *MsgSession) CreatePdr() (pdr *PDR) {
	pdr = s.upf.createPdr()
	s.pdrs = append(s.pdrs, pdr)
	return
}

func (s *MsgSession) RemovePdrs() {
	for _, pdr := range s.pdrs {
		s.upf.removePdr(pdr)
	}
	s.pdrs = []*PDR{}
}

/*
//send Session Release
func (s *MsgSession) Release() (err error) {
	return
}

//send Session Establishment/Modification (if remoteseid exists)
func (s *MsgSession) Update() (err error) {
	return
}
*/
