/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type AccNetChId struct {
	AccNetChaIdValue int32 `json:"accNetChaIdValue"`

	// Contains the identifier of the PCC rule(s) associated to the provided Access Network Charging Identifier.
	RefPccRuleIds []string `json:"refPccRuleIds,omitempty"`

	// When it is included and set to true, indicates the Access Network Charging Identifier applies to the whole PDU Session
	SessionChScope bool `json:"sessionChScope,omitempty"`
}