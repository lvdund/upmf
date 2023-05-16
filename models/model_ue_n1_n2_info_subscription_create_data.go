/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type UeN1N2InfoSubscriptionCreateData struct {
	N2InformationClass N2InformationClass `json:"n2InformationClass,omitempty"`

	N2NotifyCallbackUri string `json:"n2NotifyCallbackUri,omitempty"`

	N1MessageClass N1MessageClass `json:"n1MessageClass,omitempty"`

	N1NotifyCallbackUri string `json:"n1NotifyCallbackUri,omitempty"`

	NfId string `json:"nfId,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	OldGuami Guami `json:"oldGuami,omitempty"`
}
