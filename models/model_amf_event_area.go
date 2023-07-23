/*
Namf_EventExposure

AMF Event Exposure Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type AmfEventArea struct {
	PresenceInfo PresenceInfo `json:"presenceInfo,omitempty"`

	LadnInfo LadnInfo `json:"ladnInfo,omitempty"`

	SNssai Snssai `json:"sNssai,omitempty"`

	NsiId string `json:"nsiId,omitempty"`
}