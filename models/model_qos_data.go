/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type QosData struct {

	// Univocally identifies the QoS control policy data within a PDU session.
	QosId string `json:"qosId"`

	Var5qi int32 `json:"5qi,omitempty"`

	MaxbrUl string `json:"maxbrUl,omitempty"`

	MaxbrDl string `json:"maxbrDl,omitempty"`

	GbrUl string `json:"gbrUl,omitempty"`

	GbrDl string `json:"gbrDl,omitempty"`

	Arp Arp `json:"arp,omitempty"`

	// Indicates whether notifications are requested from 3GPP NG-RAN when the GFBR can no longer (or again) be guaranteed for a QoS Flow during the lifetime of the QoS Flow.
	Qnc bool `json:"qnc,omitempty"`

	PriorityLevel int32 `json:"priorityLevel,omitempty"`

	AverWindow int32 `json:"averWindow,omitempty"`

	MaxDataBurstVol int32 `json:"maxDataBurstVol,omitempty"`

	// Indicates whether the QoS information is reflective for the corresponding service data flow.
	ReflectiveQos bool `json:"reflectiveQos,omitempty"`

	// Indicates, by containing the same value, what PCC rules may share resource in downlink direction.
	SharingKeyDl string `json:"sharingKeyDl,omitempty"`

	// Indicates, by containing the same value, what PCC rules may share resource in uplink direction.
	SharingKeyUl string `json:"sharingKeyUl,omitempty"`

	MaxPacketLossRateDl int32 `json:"maxPacketLossRateDl,omitempty"`

	MaxPacketLossRateUl int32 `json:"maxPacketLossRateUl,omitempty"`

	// Indicates that the dynamic PCC rule shall always have its binding with the QoS Flow associated with the default QoS rule
	DefQosFlowIndication bool `json:"defQosFlowIndication,omitempty"`

	ExtMaxDataBurstVol int32 `json:"extMaxDataBurstVol,omitempty"`

	PacketDelayBudget int32 `json:"packetDelayBudget,omitempty"`

	PacketErrorRate string `json:"packetErrorRate,omitempty"`
}
