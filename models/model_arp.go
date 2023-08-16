package models

type Arp struct {

	// nullable true shall not be used for this attribute
	PriorityLevel int32 `json:"priorityLevel"`

	PreemptCap PreemptionCapability `json:"preemptCap"`

	PreemptVuln PreemptionVulnerability `json:"preemptVuln"`
}
