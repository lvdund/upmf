/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type LossOfConnectivityReason string

// List of LossOfConnectivityReason
const (
	LOSSOFCONNECTIVITYREASON_DEREGISTERED               LossOfConnectivityReason = "DEREGISTERED"
	LOSSOFCONNECTIVITYREASON_MAX_DETECTION_TIME_EXPIRED LossOfConnectivityReason = "MAX_DETECTION_TIME_EXPIRED"
	LOSSOFCONNECTIVITYREASON_PURGED                     LossOfConnectivityReason = "PURGED"
)
