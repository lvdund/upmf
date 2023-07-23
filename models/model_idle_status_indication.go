/*
Namf_EventExposure

AMF Event Exposure Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

import (
	"time"
)

type IdleStatusIndication struct {
	TimeStamp time.Time `json:"timeStamp,omitempty"`

	ActiveTime int32 `json:"activeTime,omitempty"`

	SubsRegTimer int32 `json:"subsRegTimer,omitempty"`

	EdrxCycleLength int32 `json:"edrxCycleLength,omitempty"`

	SuggestedNumOfDlPackets int32 `json:"suggestedNumOfDlPackets,omitempty"`
}