package context

import (
	"fmt"
	"net"
	"strings"
	"sync/atomic"
	"upmf/models"
)

type TopoNode struct {
	//for statemachine
	// fsm.State
	// worker    common.Executer
	// hbtimer   common.UeTimer //heartbeat timer
	// assotimer common.UeTimer //association timer
	// rectimer  common.UeTimer //recovery timer

	Heartbeat int // heartbeat interval
	Hbtry     int //num of times trying to send a heartbeat

	//attributes
	Id     string              // unique in a topo
	Infs   map[string][]NetInf //inf identities as keys
	Slices []models.Snssai

	Sbiinfo  SbiInfo
	Static   bool //is the node from config?
	Isactive uint32
}

type SbiInfo struct {
	Ip   net.IP
	Port int
}

func (sbiinfo *SbiInfo) NodeId() string {
	return fmt.Sprintf("%s:%d", sbiinfo.Ip.String(), sbiinfo.Port)
}
func (node *TopoNode) HasSbiIp() bool {
	return len(node.Sbiinfo.Ip) > 0
}

func (topoNode *TopoNode) GetTopoNodeId() string {
	return topoNode.Id
}

func (node *TopoNode) IsActive() bool {
	return atomic.LoadUint32(&node.Isactive) == 1
}

// does the node serve the snssai?
func (node *TopoNode) Serve(snssai models.Snssai) bool {
	for _, s := range node.Slices {
		if s.Sst == snssai.Sst && strings.Compare(s.Sd, snssai.Sd) == 0 {
			return true
		}
	}
	return false
}

func NewNode(id string, heartbeat int, sbiinfo *SbiConfig, static bool) (node *TopoNode) {
	node = &TopoNode{
		// State:     fsm.NewState(fsm.UPF_NONASSOCIATED),
		Id:        id,
		Heartbeat: heartbeat,
		Static:    static,
		Infs:      make(map[string][]NetInf),
		Isactive:  1,
	}
	if sbiinfo != nil {
		node.Sbiinfo.Ip = net.ParseIP(sbiinfo.Ip)
		node.Sbiinfo.Port = sbiinfo.Port
		// log.Infof("%s has pfcpinfo %s[%d]", node.id, sbiinfo.Ip, len(node.sbiinfo.Ip))
	}
	return
}

// get a NetInf from the node
func (node *TopoNode) GetInf(network string, index int) (inf *NetInf, err error) {
	if infs, ok := node.Infs[network]; ok {
		if index >= len(infs) {
			err = fmt.Errorf("Get interface at network '%s' of node '%s' error: index out of range[len=%d;index=%d]", network, node.Id, len(infs), index)
		} else {
			inf = &infs[index]
		}
	} else {
		err = fmt.Errorf("node '%s' is not in network '%s'", node.Id, network)
	}
	return
}
