package models

type PreemptionCapability string

// List of PreemptionCapability
const (
	PREEMPTIONCAPABILITY_NOT_PREEMPT PreemptionCapability = "NOT_PREEMPT"
	PREEMPTIONCAPABILITY_MAY_PREEMPT PreemptionCapability = "MAY_PREEMPT"
)