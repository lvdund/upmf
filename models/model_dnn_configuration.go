/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type DnnConfiguration struct {
	PduSessionTypes PduSessionTypes `json:"pduSessionTypes"`

	SscModes SscModes `json:"sscModes"`

	IwkEpsInd bool `json:"iwkEpsInd,omitempty"`

	Var5gQosProfile SubscribedDefaultQos `json:"5gQosProfile,omitempty"`

	SessionAmbr Ambr `json:"sessionAmbr,omitempty"`

	Var3gppChargingCharacteristics string `json:"3gppChargingCharacteristics,omitempty"`

	StaticIpAddress []IpAddress `json:"staticIpAddress,omitempty"`

	UpSecurity UpSecurity `json:"upSecurity,omitempty"`

	PduSessionContinuityInd PduSessionContinuityInd `json:"pduSessionContinuityInd,omitempty"`

	// Identity of the NEF
	NiddNefId string `json:"niddNefId,omitempty"`

	NiddInfo NiddInformation `json:"niddInfo,omitempty"`

	RedundantSessionAllowed bool `json:"redundantSessionAllowed,omitempty"`

	AcsInfo AcsInfo `json:"acsInfo,omitempty"`

	Ipv4FrameRouteList []FrameRouteInfo `json:"ipv4FrameRouteList,omitempty"`

	Ipv6FrameRouteList []FrameRouteInfo `json:"ipv6FrameRouteList,omitempty"`

	AtsssAllowed bool `json:"atsssAllowed,omitempty"`

	SecondaryAuth bool `json:"secondaryAuth,omitempty"`

	DnAaaIpAddressAllocation bool `json:"dnAaaIpAddressAllocation,omitempty"`

	DnAaaAddress IpAddress `json:"dnAaaAddress,omitempty"`

	IptvAccCtrlInfo string `json:"iptvAccCtrlInfo,omitempty"`
}
