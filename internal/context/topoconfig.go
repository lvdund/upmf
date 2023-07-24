package context

import (
	"upmf/models"
)

const (
	NET_TYPE_AN   uint8 = 0 //connect to RAN nodes
	NET_TYPE_TRAN uint8 = 1 //between two UPFs
	NET_TYPE_DNN  uint8 = 2 //UPF to DN

	SBI_DEFAULT_IP = "0.0.0.0"
)

// topo.json
type TopoConfig struct {
	Sbi      SbiConfig                `json:"sbi"`
	Networks NetConfig                `json:"networks"`
	Nodes    map[string]NodeConfig    `json:"nodes"`
	Links    map[string][]LinkConfig  `json:"links"`
	Slices   map[string]models.Snssai `json:"slices"`
}

type SbiConfig struct {
	Ip        string `json:"ip"`
	Port      int    `json:"port"`
	Heartbeat int    `json:"heartbeat,omitempty"`
}

type NetConfig struct {
	Access    []string `json:"access"`
	Transport []string `json:"transport"`
	Dnn       []string `json:"dnn"`
}

type NodeConfig struct {
	Id     string                    `json:"id"`
	Slices []string                  `json:"slices"`
	Infs   map[string][]NetInfConfig `json:"infs"`
	Sbi    *SbiConfig                `json:"sbi,omitempty"`
	Msg    string                    `json:"msg,omitempty"`
}
type NetInfConfig struct {
	Addr    string         `json:"addr"`
	DnnInfo *DnnInfoConfig `json:"dnninfo,omitempty"`
}
type DnnInfoConfig struct {
	Cidr string `json:"cidr"`
}

type LinkConfig struct {
	A LinkEndpointConfig `json:"a"`
	B LinkEndpointConfig `json:"b"`
	W *uint16            `json:"w,omitempty"`
}
type LinkEndpointConfig struct {
	Node  string `json:"node"`
	Index *int   `json:"index,omitempty"`
}
