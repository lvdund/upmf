/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type SbiBindingLevel string

// List of SbiBindingLevel
const (
	SBIBINDINGLEVEL_INSTANCE_BINDING         SbiBindingLevel = "NF_INSTANCE_BINDING"
	SBIBINDINGLEVEL_SET_BINDING              SbiBindingLevel = "NF_SET_BINDING"
	SBIBINDINGLEVEL_SERVICE_SET_BINDING      SbiBindingLevel = "NF_SERVICE_SET_BINDING"
	SBIBINDINGLEVEL_SERVICE_INSTANCE_BINDING SbiBindingLevel = "NF_SERVICE_INSTANCE_BINDING"
)