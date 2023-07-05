package context

import (
	"upmf/models"
)

// For create upmf
type UPMF struct {
	UpfTopo *UpfTopo `json:"upftopo"`
	// Nodes   map[string]*upman.UpNode `json:"nodes"`
	Config  *UpmfConfig           `json:"config"`
	ListMap map[string][]QueryMap `json:"map"`
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