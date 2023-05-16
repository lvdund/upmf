/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type BatteryIndicationRm struct {
	BatteryInd bool `json:"batteryInd,omitempty"`

	ReplaceableInd bool `json:"replaceableInd,omitempty"`

	RechargeableInd bool `json:"rechargeableInd,omitempty"`
}
