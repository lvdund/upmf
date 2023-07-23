/*
Namf_Location

AMF Location Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type PositioningMethod string

// List of PositioningMethod
const (
	POSITIONINGMETHOD_CELLID              PositioningMethod = "CELLID"
	POSITIONINGMETHOD_ECID                PositioningMethod = "ECID"
	POSITIONINGMETHOD_OTDOA               PositioningMethod = "OTDOA"
	POSITIONINGMETHOD_BAROMETRIC_PRESSURE PositioningMethod = "BAROMETRIC_PRESSURE"
	POSITIONINGMETHOD_WLAN                PositioningMethod = "WLAN"
	POSITIONINGMETHOD_BLUETOOTH           PositioningMethod = "BLUETOOTH"
	POSITIONINGMETHOD_MBS                 PositioningMethod = "MBS"
	POSITIONINGMETHOD_MOTION_SENSOR       PositioningMethod = "MOTION_SENSOR"
	POSITIONINGMETHOD_DL_TDOA             PositioningMethod = "DL_TDOA"
	POSITIONINGMETHOD_DL_AOD              PositioningMethod = "DL_AOD"
	POSITIONINGMETHOD_MULTI_RTT           PositioningMethod = "MULTI-RTT"
	POSITIONINGMETHOD_NR_ECID             PositioningMethod = "NR_ECID"
	POSITIONINGMETHOD_UL_TDOA             PositioningMethod = "UL_TDOA"
	POSITIONINGMETHOD_UL_AOA              PositioningMethod = "UL_AOA"
	POSITIONINGMETHOD_NETWORK_SPECIFIC    PositioningMethod = "NETWORK_SPECIFIC"
)