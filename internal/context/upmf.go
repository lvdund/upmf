package context

import (
	"upmf/models"
)

// For create upmf
type UPMF struct {
	UpfTopo      *UpfTopo                   `json:"upftopo"`
	ListLinks    map[models.Snssai][]Link `json:"listlink"`
	Config       *UpmfConfig                `json:"config"`
	ListQueryMap map[string][]QueryMap      `json:"map"`
	// Nodes   map[string]*upman.UpNode `json:"nodes"`
	// Pfcp    *pfcp.Pfcp               `json:"pfcp`
}

type UpmfConfig struct {
	NfInstanceID string                   `json:"nfInstanceid,omitempty"`
	Name         string                   `json:"name"`
	Sbi          SbiInfo                  `json:"sbi"`
	Networks     NetConfig                `json:"network,omitempty"`
	PlmnList     *[]models.PlmnId         `json:"plmnlist"`
	Slices       map[string]models.Snssai `json:"slices"`
}
