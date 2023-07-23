/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type MonitoringConfiguration struct {
	EventType EventType `json:"eventType"`

	ImmediateFlag bool `json:"immediateFlag,omitempty"`

	LocationReportingConfiguration LocationReportingConfiguration `json:"locationReportingConfiguration,omitempty"`

	AssociationType AssociationType `json:"associationType,omitempty"`

	DatalinkReportCfg DatalinkReportingConfiguration `json:"datalinkReportCfg,omitempty"`

	LossConnectivityCfg LossConnectivityCfg `json:"lossConnectivityCfg,omitempty"`

	MaximumLatency int32 `json:"maximumLatency,omitempty"`

	MaximumResponseTime int32 `json:"maximumResponseTime,omitempty"`

	SuggestedPacketNumDl int32 `json:"suggestedPacketNumDl,omitempty"`

	PduSessionStatusCfg PduSessionStatusCfg `json:"pduSessionStatusCfg,omitempty"`

	ReachabilityForSmsCfg ReachabilityForSmsConfiguration `json:"reachabilityForSmsCfg,omitempty"`

	MtcProviderInformation string `json:"mtcProviderInformation,omitempty"`

	AfId string `json:"afId,omitempty"`

	IdleStatusInd bool `json:"idleStatusInd,omitempty"`
}