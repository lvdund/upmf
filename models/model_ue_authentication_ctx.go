/*
AUSF API

AUSF UE Authentication Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type UEAuthenticationCtx struct {
	AuthType AuthType `json:"authType"`

	Var5gAuthData UEAuthenticationCtx5gAuthData `json:"5gAuthData"`
	
	EAPMessage	string `json:"EapPayload"`

	Links map[string]LinksValueSchema `json:"_links"`

	ServingNetworkName string `json:"servingNetworkName,omitempty"`
}
