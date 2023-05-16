package context

import (
	"strings"
	"upmf/models"
)

type UpfTopo struct {
	Nets       map[string]uint8
	Nodes      map[string]*TopoNode
	Links      []Link
	Sbiid2node map[string]*TopoNode

	Heartbeat int
}

func (topo *UpfTopo) GetNodes() map[string]*TopoNode {
	return topo.Nodes
}

func (topo *UpfTopo) GetHeartbeat() int {
	return topo.Heartbeat
}

// a link between two network interfaces
type Link struct {
	Inf1 *NetInf
	Inf2 *NetInf
	W    uint16
}

// is this link active and can it serve the given slice?
func (l *Link) IsActive(snssai models.Snssai) bool {
	return l.Inf1.Local.IsActive() && l.Inf2.Local.IsActive() &&
		l.Inf1.Local.Serve(snssai) && l.Inf2.Local.Serve(snssai)
}

func CreatenewTopo() *UpfTopo {
	return &UpfTopo{
		Nets:       make(map[string]uint8),
		Nodes:      make(map[string]*TopoNode),
		Sbiid2node: make(map[string]*TopoNode),
	}
}

// Get a node's network interfaces to Access networks
func (topo *UpfTopo) GetNodeAnFaces(node *TopoNode, nets []string) (foundinfs []NetInf) {
	for network, infs := range node.Infs {
		if ntype, ok := topo.Nets[network]; ok && ntype == NET_TYPE_AN {
			for _, netname := range nets {
				if strings.Compare(netname, network) == 0 {
					foundinfs = append(foundinfs, infs...)
					break
				}
			}
		}
	}
	return
}

// Get a node's network interfaces to Dnn
func (topo *UpfTopo) GetNodeDnnFaces(node *TopoNode, dnn string) (foundinfs []NetInf) {
	for network, infs := range node.Infs {
		if ntype, ok := topo.Nets[network]; ok && ntype == NET_TYPE_DNN && strings.Compare(network, dnn) == 0 {
			foundinfs = infs
			break
		}
	}
	return
}
