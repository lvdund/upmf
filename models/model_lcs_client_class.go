/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type LcsClientClass string

// List of LcsClientClass
const (
	LCSCLIENTCLASS_BROADCAST_SERVICE          LcsClientClass = "BROADCAST_SERVICE"
	LCSCLIENTCLASS_OM_IN_HPLMN                LcsClientClass = "OM_IN_HPLMN"
	LCSCLIENTCLASS_OM_IN_VPLMN                LcsClientClass = "OM_IN_VPLMN"
	LCSCLIENTCLASS_ANONYMOUS_LOCATION_SERVICE LcsClientClass = "ANONYMOUS_LOCATION_SERVICE"
	LCSCLIENTCLASS_SPECIFIC_SERVICE           LcsClientClass = "SPECIFIC_SERVICE"
)