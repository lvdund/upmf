/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// SmPolicySnssaiDataPatch - Contains the SM policy data for a given subscriber and S-NSSAI.
type SmPolicySnssaiDataPatch struct {
	Snssai Snssai `json:"snssai"`

	SmPolicyDnnData map[string]SmPolicyDnnDataPatch `json:"smPolicyDnnData,omitempty"`
}
