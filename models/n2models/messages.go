package n2models

import "upmf/models"

const (
	N2MSG_NAS_INIT uint8 = iota
	N2MSG_NAS_UL
	N2MSG_NAS_DL
	N2MSG_NAS_NON_DLVR

	N2MSG_UECTX_REL_REQ
	N2MSG_UECTX_REL_CMPL
	N2MSG_UECTX_MOD_RSP
	N2MSG_UECTX_MOD_FAIL

	N2MSG_RRC_INACT_TRAN_REP

	N2MSG_UECTX_SET_REQ
	N2MSG_UECTX_SET_RSP
	N2MSG_UECTX_SET_FAIL

	N2MSG_PDU_SET_RSP
	N2MSG_PDU_REL_RSP
	N2MSG_PDU_MOD_RSP
	N2MSG_PDU_NOT
	N2MSG_PDU_MOD_IND

	N2MSG_PDU_SET_REQ
)

type N2Msg struct {
	Code    uint8
	Content interface{}
}

type InitialUeMsg struct {
	NasPdu         []byte
	ContextRequest bool
	RrcCause       uint8
	Loc            *models.UserLocation
}

// Uplink Nas Transport
type UlNasTransport struct {
	NasPdu []byte
	Loc    *models.UserLocation
}

type NasNonDeliveryIndication struct {
	NasPdu []byte
	Cause  Cause
}

// Downlink Nas Transport
type NasDlMsg struct {
	NasPdu           []byte
	MobiRestrictList *MobilityRestrictionList //optional
	UeAmbr           *UeAmbr                  //optional
	AllowedNssai     *models.AllowedNssai     //optional
	OldAmf           string                   //optional
}

// Paging
type PagingMsg struct {
	Non3Gpp bool
	PP      PagingPriority
}
