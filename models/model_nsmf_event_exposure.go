/*
Nsmf_EventExposure

Session Management Event Exposure Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

import (
	"time"
)

// NsmfEventExposure - Represents an Individual SMF Notification Subscription resource. The serviveName property corresponds to the serviceName in the main body of the specification.
type NsmfEventExposure struct {
	Supi string `json:"supi,omitempty"`

	Gpsi string `json:"gpsi,omitempty"`

	// Any UE indication. This IE shall be present if the event subscription is applicable to any UE. Default value \"false\" is used, if not present.
	AnyUeInd bool `json:"anyUeInd,omitempty"`

	GroupId string `json:"groupId,omitempty"`

	PduSeId int32 `json:"pduSeId,omitempty"`

	Dnn string `json:"dnn,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	// Identifies an Individual SMF Notification Subscription. To enable that the value is used as part of a URI, the string shall only contain characters allowed according to the \"lower-with-hyphen\" naming convention defined in 3GPP TS 29.501. In an OpenAPI schema, the format shall be designated as \"SubId\".
	SubId string `json:"subId,omitempty"`

	// Notification Correlation ID assigned by the NF service consumer.
	NotifId string `json:"notifId"`

	NotifUri string `json:"notifUri"`

	// Alternate or backup IPv4 address(es) where to send Notifications.
	AltNotifIpv4Addrs []string `json:"altNotifIpv4Addrs,omitempty"`

	// Alternate or backup IPv6 address(es) where to send Notifications.
	AltNotifIpv6Addrs []Ipv6Addr `json:"altNotifIpv6Addrs,omitempty"`

	// Alternate or backup FQDN(s) where to send Notifications.
	AltNotifFqdns []string `json:"altNotifFqdns,omitempty"`

	// Subscribed events
	EventSubs []EventSubscription `json:"eventSubs"`

	ImmeRep bool `json:"ImmeRep,omitempty"`

	NotifMethod NotificationMethod `json:"notifMethod,omitempty"`

	MaxReportNbr int32 `json:"maxReportNbr,omitempty"`

	Expiry time.Time `json:"expiry,omitempty"`

	RepPeriod int32 `json:"repPeriod,omitempty"`

	Guami Guami `json:"guami,omitempty"`

	ServiveName ServiceName `json:"serviveName,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	SampRatio int32 `json:"sampRatio,omitempty"`

	GrpRepTime int32 `json:"grpRepTime,omitempty"`
}