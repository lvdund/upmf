/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type NonUeN2InfoSubscriptionCreateData struct {
	GlobalRanNodeList []GlobalRanNodeId `json:"globalRanNodeList,omitempty"`

	AnTypeList []AccessType `json:"anTypeList,omitempty"`

	N2InformationClass N2InformationClass `json:"n2InformationClass"`

	N2NotifyCallbackUri string `json:"n2NotifyCallbackUri"`

	NfId string `json:"nfId,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}