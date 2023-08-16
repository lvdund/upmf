package models

type AuthorizedDefaultQos struct {
	Var5qi int32 `json:"5qi,omitempty"`

	Arp Arp `json:"arp,omitempty"`

	PriorityLevel int32 `json:"priorityLevel,omitempty"`

	AverWindow int32 `json:"averWindow,omitempty"`

	MaxDataBurstVol int32 `json:"maxDataBurstVol,omitempty"`

	MaxbrUl string `json:"maxbrUl,omitempty"`

	MaxbrDl string `json:"maxbrDl,omitempty"`

	GbrUl string `json:"gbrUl,omitempty"`

	GbrDl string `json:"gbrDl,omitempty"`

	ExtMaxDataBurstVol int32 `json:"extMaxDataBurstVol,omitempty"`
}
