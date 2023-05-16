/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type PresenceState string

// List of PresenceState
const (
	PRESENCESTATE_IN_AREA     PresenceState = "IN_AREA"
	PRESENCESTATE_OUT_OF_AREA PresenceState = "OUT_OF_AREA"
	PRESENCESTATE_UNKNOWN     PresenceState = "UNKNOWN"
	PRESENCESTATE_INACTIVE    PresenceState = "INACTIVE"
)
