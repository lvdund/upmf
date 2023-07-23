/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type RequestedUsageData struct {

	// An array of usage monitoring data id references to the usage monitoring data instances for which the PCF is requesting a usage report. This attribute shall only be provided when allUmIds is not set to true.
	RefUmIds []string `json:"refUmIds,omitempty"`

	// This boolean indicates whether requested usage data applies to all usage monitoring data instances. When it's not included, it means requested usage data shall only apply to the usage monitoring data instances referenced by the refUmIds attribute.
	AllUmIds bool `json:"allUmIds,omitempty"`
}