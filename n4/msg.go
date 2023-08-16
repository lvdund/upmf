package n4

import (
	n4msg "github.com/lvdund/n4interface"
	msgType "github.com/lvdund/n4interface/msgType"
)

func (session *MsgSession) FillEstablishmentRequest(msg *n4msg.MSGSessionEstablishmentRequest) {
	log.Trace("Fill MSGSessionEstablishmentRequest")
	isv4 := true

	msg.CPFSEID = &msgType.FSEID{
		V4:          isv4,
		V6:          !isv4,
		Seid:        session.localseid,
		Ipv4Address: session.upf.ip,
	}

	pdrlist, farlist, barlist, qerlist, urrlist := getRuleList(session.pdrs, []RuleState{RULE_INITIAL})

	for _, pdr := range pdrlist {
		msg.CreatePDR = append(msg.CreatePDR, pdr.toCreatePdr())
		pdr.State = RULE_CREATE
	}  

	for _, far := range farlist {
		msg.CreateFAR = append(msg.CreateFAR, far.toCreateFar())
		far.State = RULE_CREATE
	}

	for _, bar := range barlist {
		msg.CreateBAR = append(msg.CreateBAR, bar.toCreateBar())
		bar.State = RULE_CREATE
	}

	for _, qer := range qerlist {
		msg.CreateQER = append(msg.CreateQER, qer.toCreateQer())
		qer.State = RULE_CREATE
	}
	for _, urr := range urrlist {
		msg.CreateURR = append(msg.CreateURR, urr.toCreateUrr())
		urr.State = RULE_CREATE
	}

	msg.PDNType = &msgType.PDNType{
		PdnType: msgType.PDNTypeIpv4,
	}
	return
}

//get rule lists for create/update/delete
func getRuleList(pdrs []*PDR, states []RuleState) (pdrlist []*PDR, farlist []*FAR, barlist []*BAR, qerlist []*QER, urrlist []*URR) {
	qers := make(map[uint32]bool)
	urrs := make(map[uint32]bool)
	for _, pdr := range pdrs {
		if inState(pdr.State, states) {
			pdrlist = append(pdrlist, pdr)
		}
		if far := pdr.FAR; far != nil {
			if inState(far.State, states) {
				farlist = append(farlist, far)
			}
			if bar := far.BAR; bar != nil && inState(bar.State, states) {
				barlist = append(barlist, bar)
			}
		}

		for _, urr := range pdr.URR {
			if _, ok := urrs[urr.URRID]; !ok && inState(urr.State, states) {
				urrlist = append(urrlist, urr)
				urrs[urr.URRID] = true
			}
		}
		for _, qer := range pdr.QER {
			if _, ok := qers[qer.QERID]; !ok && inState(qer.State, states) {
				qerlist = append(qerlist, qer)
				qers[qer.QERID] = true
			}
		}
	}
	return
}

func inState(s RuleState, states []RuleState) bool {
	for _, state := range states {
		if s == state {
			return true
		}
	}
	return false
}