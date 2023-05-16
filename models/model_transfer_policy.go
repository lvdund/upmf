/*
Npcf_BDTPolicyControl Service API

PCF BDT Policy Control Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// TransferPolicy - Describes a transfer policy.
type TransferPolicy struct {
	MaxBitRateDl string `json:"maxBitRateDl,omitempty"`

	MaxBitRateUl string `json:"maxBitRateUl,omitempty"`

	// Indicates a rating group for the recommended time window.
	RatingGroup int32 `json:"ratingGroup"`

	RecTimeInt TimeWindow `json:"recTimeInt"`

	// Contains an identity of a transfer policy.
	TransPolicyId int32 `json:"transPolicyId"`
}
