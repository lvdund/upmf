/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type NonUeN2InfoSubscriptionCreatedData struct {
	N2NotifySubscriptionId string `json:"n2NotifySubscriptionId"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	N2InformationClass N2InformationClass `json:"n2InformationClass,omitempty"`
}