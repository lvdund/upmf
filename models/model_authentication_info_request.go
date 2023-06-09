/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type AuthenticationInfoRequest struct {
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	ServingNetworkName string `json:"servingNetworkName"`

	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`

	AusfInstanceId string `json:"ausfInstanceId"`

	CellCagInfo []string `json:"cellCagInfo,omitempty"`

	N5gcInd bool `json:"n5gcInd,omitempty"`
}
