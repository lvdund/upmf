/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type PpProfileData struct {

	// A map (list of key-value pairs where PpDataType serves as key) of AllowedMtcProviderInfo lists. In addition to defined PpDataType, the key value \"ALL\" may be used to identify a map entry which contains a list of AllowedMtcProviderInfo that are allowed to provision all types of the PP data for the user using UDM ParameterProvision service.
	AllowedMtcProviders map[string][]AllowedMtcProviderInfo `json:"allowedMtcProviders,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}