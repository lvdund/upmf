/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type SmContextReleaseData struct {
	Cause Cause `json:"cause,omitempty"`

	NgApCause NgApCause `json:"ngApCause,omitempty"`

	Var5gMmCauseValue int32 `json:"5gMmCauseValue,omitempty"`

	UeLocation UserLocation `json:"ueLocation,omitempty"`

	UeTimeZone string `json:"ueTimeZone,omitempty"`

	AddUeLocation UserLocation `json:"addUeLocation,omitempty"`

	VsmfReleaseOnly bool `json:"vsmfReleaseOnly,omitempty"`

	N2SmInfo RefToBinaryData `json:"n2SmInfo,omitempty"`

	N2SmInfoType N2SmInfoType `json:"n2SmInfoType,omitempty"`

	IsmfReleaseOnly bool `json:"ismfReleaseOnly,omitempty"`
}
