package models

type UpfConfig struct {
	Id     string                    `json:"id"`
	Slices []string                  `json:"slices"`
	Infs   map[string][]NetInfConfig `json:"infs"`
	Sbi    Sbi                       `json:"sbi,omitempty"`
}

type NetInfConfig struct {
	NetName string         `json:"netname,omitempty"`
	Addr    string         `json:"addr"`
	DnnInfo *DnnInfoConfig `json:"dnninfo,omitempty"`
}

type DnnInfoConfig struct {
	Cidr string `json:"cidr"`
}
