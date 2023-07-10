package models

import (
	"net"
	"strings"
)

type Query struct {
	SmfId string    `json:"smfid"`
	UeId  string    `json:"ueid"`
	Query PathQuery `json:"query`
}

// represent a query to search for a data path to serve a given pdu session
type PathQuery struct {
	Dnn    string   `json:"dnn"`    //to locate anchor UPF
	Snssai Snssai   `json:"snssai"` //all UPFs must serve this snssai
	Nets   []string `json:"nets"`   //names of UP networks that connect to RAN UP (to find the first RAN-connected UPF)
}

type DataPath struct {
	Path []*PathNode `json:"path"` //UP nodes in the path
	Ip   net.IP      `json:"ip"`   //Ue IP allocated by the anchor node
	// Deallocator func(net.IP) `json:"deallocator,omitempty"` //for release Ue IP
}

type PathNode struct {
	Id       string //Upf Identity (read from its configuration)
	UlIp     net.IP //ip of the uplink face
	DlIp     net.IP //ip of the downlink face
	PfcpIp   net.IP
	PfcpPort int
}

func (p *DataPath) String() string {
	nodes := []string{}
	for _, n := range p.Path {
		nodes = append(nodes, n.String())
	}
	return strings.Join(nodes, " ")
}
func (n *PathNode) String() string {
	//return fmt.Sprintf("%s_%s_%s", n.dlip, n.id, n.ulip)
	return n.Id
}
func (p *DataPath) ReleaseIp() {
	// p.Deallocator(p.Ip)
}
