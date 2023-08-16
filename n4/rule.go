package n4

import (
	msg "github.com/lvdund/n4interface"
	msgType "github.com/lvdund/n4interface/msgType"
)

const (
	RULE_INITIAL RuleState = 0
	RULE_CREATE  RuleState = 1
	RULE_UPDATE  RuleState = 2
	RULE_REMOVE  RuleState = 3
)

type RuleState uint8

// Packet Detection Rule. Table 7.5.2.2-1
type PDR struct {
	PDRID uint16

	Precedence         uint32
	PDI                PDI
	OuterHeaderRemoval *msgType.OuterHeaderRemoval

	FAR *FAR
	URR []*URR
	QER []*QER

	State RuleState
}

func (pdr *PDR) toCreatePdr() *msg.CreatePDR {
	cpdr := &msg.CreatePDR{
		PDRID: &msgType.PacketDetectionRuleID{
			RuleId: pdr.PDRID,
		},
		Precedence: &msgType.Precedence{
			PrecedenceValue: pdr.Precedence,
		},
		PDI: &msg.PDI{
			SourceInterface: &pdr.PDI.SourceInterface,
			LocalFTEID:      pdr.PDI.LocalFTeid,
			NetworkInstance: pdr.PDI.NetworkInstance,
			UEIPAddress:     pdr.PDI.UEIPAddress,
		},
		OuterHeaderRemoval: pdr.OuterHeaderRemoval,
	}

	if pdr.PDI.ApplicationID != "" {
		cpdr.PDI.ApplicationID = &msgType.ApplicationID{
			ApplicationIdentifier: []byte(pdr.PDI.ApplicationID),
		}
	}

	if pdr.PDI.SDFFilter != nil {
		cpdr.PDI.SDFFilter = pdr.PDI.SDFFilter
	}

	if far := pdr.FAR; far != nil {
		cpdr.FARID = &msgType.FARID{
			FarIdValue: far.FARID,
		}
	}
	for _, urr := range pdr.URR {
		cpdr.URRID = append(cpdr.URRID, &msgType.URRID{
			UrrIdValue: urr.URRID,
		})
	}

	for _, qer := range pdr.QER {
		cpdr.QERID = append(cpdr.QERID, &msgType.QERID{
			QERID: qer.QERID,
		})
	}

	return cpdr
}

func (pdr *PDR) toUpdatePdr() *msg.UpdatePDR {
	updr := &msg.UpdatePDR{
		PDRID: &msgType.PacketDetectionRuleID{
			RuleId: pdr.PDRID,
		},
		Precedence: &msgType.Precedence{
			PrecedenceValue: pdr.Precedence,
		},
		PDI: &msg.PDI{
			SourceInterface: &pdr.PDI.SourceInterface,
			LocalFTEID:      pdr.PDI.LocalFTeid,
			NetworkInstance: pdr.PDI.NetworkInstance,
			UEIPAddress:     pdr.PDI.UEIPAddress,
		},
	}

	if pdr.PDI.ApplicationID != "" {
		updr.PDI.ApplicationID = &msgType.ApplicationID{
			ApplicationIdentifier: []byte(pdr.PDI.ApplicationID),
		}
	}

	if pdr.PDI.SDFFilter != nil {
		updr.PDI.SDFFilter = pdr.PDI.SDFFilter
	}

	updr.OuterHeaderRemoval = pdr.OuterHeaderRemoval

	if pdr.FAR != nil {
		updr.FARID = &msgType.FARID{
			FarIdValue: pdr.FAR.FARID,
		}
	}
	return updr
}

func (pdr *PDR) toRemovePdr() *msg.RemovePDR {
	return &msg.RemovePDR{
		PDRID: &msgType.PacketDetectionRuleID{
			RuleId: pdr.PDRID,
		},
	}
}

// Packet Detection.
// 7.5.2.2-2
type PDI struct {
	SourceInterface msgType.SourceInterface
	LocalFTeid      *msgType.FTEID
	NetworkInstance *msgType.NetworkInstance
	UEIPAddress     *msgType.UEIPAddress
	SDFFilter       *msgType.SDFFilter
	ApplicationID   string
}

// Forwarding Action Rule
type FAR struct {
	FARID uint32

	ApplyAction          msgType.ApplyAction
	ForwardingParameters *ForwardingParameters

	BAR   *BAR
	State RuleState
}

func (far *FAR) toCreateFar() *msg.CreateFAR {
	//add createFAR
	cfar := &msg.CreateFAR{
		FARID: &msgType.FARID{
			FarIdValue: far.FARID,
		},
		ApplyAction: &msgType.ApplyAction{},
	}

	if far.ForwardingParameters != nil {
		cfar.ApplyAction.Forw = true
	} else {
		//	29.244 v15.3 Table 7.5.2.3-1 Farwarding Parameters IE shall be
		//	present when the Apply-Action requests the packets to be forwarded.
		//	FAR without Farwarding Parameters set Apply Action as Drop instead
		//	of Forward.

		cfar.ApplyAction.Forw = false
		cfar.ApplyAction.Drop = true
	}

	if far.ForwardingParameters != nil {
		cfar.ForwardingParameters = &msg.ForwardingParametersIEInFAR{
			DestinationInterface: &far.ForwardingParameters.DestinationInterface,
			NetworkInstance:      far.ForwardingParameters.NetworkInstance,
			OuterHeaderCreation:  far.ForwardingParameters.OuterHeaderCreation,
		}
		if far.ForwardingParameters.ForwardingPolicyID != "" {
			cfar.ForwardingParameters.ForwardingPolicy = &msgType.ForwardingPolicy{
				ForwardingPolicyIdentifierLength: uint8(len(far.ForwardingParameters.ForwardingPolicyID)),
				ForwardingPolicyIdentifier:       []byte(far.ForwardingParameters.ForwardingPolicyID),
			}
		}
	}
	if far.BAR != nil {
		cfar.BARID = &msgType.BARID{
			BarIdValue: far.BAR.BARID,
		}
	}
	return cfar
}

func (far *FAR) toUpdateFar() *msg.UpdateFAR {
	ufar := &msg.UpdateFAR{
		FARID: &msgType.FARID{
			FarIdValue: far.FARID,
		},
		ApplyAction: &msgType.ApplyAction{
			Forw: far.ApplyAction.Forw,
			Buff: far.ApplyAction.Buff,
			Nocp: far.ApplyAction.Nocp,
			Dupl: far.ApplyAction.Dupl,
			Drop: far.ApplyAction.Drop,
		},
	}
	if far.BAR != nil {
		ufar.BARID = &msgType.BARID{
			BarIdValue: far.BAR.BARID,
		}
	}

	if far.ForwardingParameters != nil {
		ufar.UpdateForwardingParameters = &msg.UpdateForwardingParametersIEInFAR{
			DestinationInterface: &far.ForwardingParameters.DestinationInterface,
			NetworkInstance:      far.ForwardingParameters.NetworkInstance,
			OuterHeaderCreation:  far.ForwardingParameters.OuterHeaderCreation,
			MSGSMReqFlags: &msgType.MSGSMReqFlags{
				Sndem: far.ForwardingParameters.SendEndMarker,
			},
		}
		if far.ForwardingParameters.ForwardingPolicyID != "" {
			ufar.UpdateForwardingParameters.ForwardingPolicy = &msgType.ForwardingPolicy{
				ForwardingPolicyIdentifierLength: uint8(len(far.ForwardingParameters.ForwardingPolicyID)),
				ForwardingPolicyIdentifier:       []byte(far.ForwardingParameters.ForwardingPolicyID),
			}
		}
	}

	return ufar
}

func (far *FAR) toRemoveFar() *msg.RemoveFAR {

	return &msg.RemoveFAR{
		FARID: &msgType.FARID{
			FarIdValue: far.FARID,
		},
	}
}

// Forwarding Parameters.
type ForwardingParameters struct {
	DestinationInterface msgType.DestinationInterface
	NetworkInstance      *msgType.NetworkInstance
	OuterHeaderCreation  *msgType.OuterHeaderCreation
	ForwardingPolicyID   string
	SendEndMarker        bool
}

// Buffering Action Rule
type BAR struct {
	BARID uint8

	DownlinkDataNotificationDelay  msgType.DownlinkDataNotificationDelay
	SuggestedBufferingPacketsCount msgType.SuggestedBufferingPacketsCount

	State RuleState
}

func (bar *BAR) toCreateBar() *msg.CreateBAR {
	return &msg.CreateBAR{
		BARID: &msgType.BARID{
			BarIdValue: bar.BARID,
		},
		//DownlinkDataNotificationDelay:  &msgType.DownlinkDataNotificationDelay{},
		//SuggestedBufferingPacketsCount: &msgType.SuggestedBufferingPacketsCount{},
	}
}

func (bar *BAR) toUpdateBar() *msg.UpdateBARMSGSessionModificationRequest {
	return nil
}

func (bar *BAR) toRemoveBar() *msg.RemoveBAR {
	return &msg.RemoveBAR{
		BARID: &msgType.BARID{
			BarIdValue: bar.BARID,
		},
	}
}

// QoS Enhancement Rule
type QER struct {
	QERID uint32

	QFI msgType.QFI

	GateStatus *msgType.GateStatus
	MBR        *msgType.MBR
	GBR        *msgType.GBR

	State RuleState
}

func (qer *QER) toCreateQer() *msg.CreateQER {
	return &msg.CreateQER{
		QERID: &msgType.QERID{
			QERID: qer.QERID,
		},
		GateStatus:        qer.GateStatus,
		QoSFlowIdentifier: &qer.QFI,
		MaximumBitrate:    qer.MBR,
		GuaranteedBitrate: qer.GBR,
	}
}

func (qer *QER) toUpdateQer() *msg.UpdateQER {
	return nil
}

func (qer *QER) toRemoveQer() *msg.RemoveQER {
	return &msg.RemoveQER{
		QERID: &msgType.QERID{
			QERID: qer.QERID,
		},
	}
}

// Usage Report Rule
type URR struct {
	URRID uint32
	//TODO: Add more atributes
	State RuleState
}

func (urr *URR) toCreateUrr() *msg.CreateURR {
	return &msg.CreateURR{
		URRID: &msgType.URRID{
			UrrIdValue: urr.URRID,
		},
	}
}

func (urr *URR) toUpdateUrr() *msg.UpdateURR {
	return &msg.UpdateURR{
		URRID: &msgType.URRID{
			UrrIdValue: urr.URRID,
		},
		//add more attributes
	}
}

func (urr *URR) toRemoveUrr() *msg.RemoveURR {
	return &msg.RemoveURR{
		URRID: &msgType.URRID{
			UrrIdValue: urr.URRID,
		},
	}
}
