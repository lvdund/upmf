package context

import (
	"upmf/models"
)

// For create upmf
type UPMF struct {
	// UpfTopo      *UpfTopo                   `json:"upftopo"`
	// ListLinks    map[models.Snssai][]Link   `json:"listlink"`
	Config       *UpmfConfig                `json:"config"`
	// ListQueryMap map[string][]QueryMap      `json:"map"`
	TopoMaps     map[models.Snssai]*UpfTopo `json:"topomaps"`
}

type UpmfConfig struct {
	NfInstanceID string                   `json:"nfInstanceid,omitempty"`
	Name         string                   `json:"name"`
	Sbi          SbiInfo                  `json:"sbi"`
	PlmnList     *[]models.PlmnId         `json:"plmnlist"`
	Slices       map[string]models.Snssai `json:"slices"`
	Nets         map[string]uint8         `json:"nets"`
}
