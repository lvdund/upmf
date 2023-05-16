/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type ProblemDetails2 struct {

	// string providing an URI formatted according to IETF RFC 3986.
	Type string `json:"type,omitempty"`

	// A short, human-readable summary of the problem type. It should not change from occurrence to occurrence of the problem.
	Title string `json:"title,omitempty"`

	// The HTTP status code for this occurrence of the problem.
	Status int32 `json:"status,omitempty"`

	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty"`

	// string providing an URI formatted according to IETF RFC 3986.
	Instance string `json:"instance,omitempty"`

	// A machine-readable application error cause specific to this occurrence of the problem. This IE should be present and provide application-related error information, if available.
	Cause string `json:"cause,omitempty"`

	// Description of invalid parameters, for a request rejected due to invalid parameters.
	InvalidParams []InvalidParam1 `json:"invalidParams,omitempty"`
}
