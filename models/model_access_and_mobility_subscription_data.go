/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type AccessAndMobilitySubscriptionData struct {
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	Gpsis []string `json:"gpsis,omitempty"`

	InternalGroupIds []string `json:"internalGroupIds,omitempty"`

	SharedVnGroupDataIds map[string]string `json:"sharedVnGroupDataIds,omitempty"`

	SubscribedUeAmbr AmbrRm `json:"subscribedUeAmbr,omitempty"`

	Nssai *Nssai `json:"nssai,omitempty"`

	RatRestrictions []RatType `json:"ratRestrictions,omitempty"`

	ForbiddenAreas []Area `json:"forbiddenAreas,omitempty"`

	ServiceAreaRestriction ServiceAreaRestriction `json:"serviceAreaRestriction,omitempty"`

	CoreNetworkTypeRestrictions []CoreNetworkType `json:"coreNetworkTypeRestrictions,omitempty"`

	RfspIndex *int32 `json:"rfspIndex,omitempty"`

	SubsRegTimer *int32 `json:"subsRegTimer,omitempty"`

	UeUsageType int32 `json:"ueUsageType,omitempty"`

	MpsPriority bool `json:"mpsPriority,omitempty"`

	McsPriority bool `json:"mcsPriority,omitempty"`

	ActiveTime *int32 `json:"activeTime,omitempty"`

	SorInfo SorInfo `json:"sorInfo,omitempty"`

	SorInfoExpectInd bool `json:"sorInfoExpectInd,omitempty"`

	SorafRetrieval bool `json:"sorafRetrieval,omitempty"`

	SorUpdateIndicatorList []SorUpdateIndicator `json:"sorUpdateIndicatorList,omitempty"`

	UpuInfo UpuInfo `json:"upuInfo,omitempty"`

	MicoAllowed bool `json:"micoAllowed,omitempty"`

	SharedAmDataIds []string `json:"sharedAmDataIds,omitempty"`

	OdbPacketServices OdbPacketServices `json:"odbPacketServices,omitempty"`

	SubscribedDnnList []AccessAndMobilitySubscriptionDataSubscribedDnnListInner `json:"subscribedDnnList,omitempty"`

	ServiceGapTime int32 `json:"serviceGapTime,omitempty"`

	MdtUserConsent MdtUserConsent `json:"mdtUserConsent,omitempty"`

	MdtConfiguration MdtConfiguration `json:"mdtConfiguration,omitempty"`

	TraceData *TraceData `json:"traceData,omitempty"`

	CagData CagData `json:"cagData,omitempty"`

	StnSr string `json:"stnSr,omitempty"`

	CMsisdn string `json:"cMsisdn,omitempty"`

	NbIoTUePriority int32 `json:"nbIoTUePriority,omitempty"`

	NssaiInclusionAllowed bool `json:"nssaiInclusionAllowed,omitempty"`

	RgWirelineCharacteristics string `json:"rgWirelineCharacteristics,omitempty"`

	EcRestrictionDataWb EcRestrictionDataWb `json:"ecRestrictionDataWb,omitempty"`

	EcRestrictionDataNb bool `json:"ecRestrictionDataNb,omitempty"`

	ExpectedUeBehaviourList ExpectedUeBehaviourData `json:"expectedUeBehaviourList,omitempty"`

	PrimaryRatRestrictions []RatType `json:"primaryRatRestrictions,omitempty"`

	SecondaryRatRestrictions []RatType `json:"secondaryRatRestrictions,omitempty"`

	EdrxParametersList []EdrxParameters `json:"edrxParametersList,omitempty"`

	PtwParametersList []PtwParameters `json:"ptwParametersList,omitempty"`

	IabOperationAllowed bool `json:"iabOperationAllowed,omitempty"`

	WirelineForbiddenAreas []WirelineArea `json:"wirelineForbiddenAreas,omitempty"`

	WirelineServiceAreaRestriction WirelineServiceAreaRestriction `json:"wirelineServiceAreaRestriction,omitempty"`
}
