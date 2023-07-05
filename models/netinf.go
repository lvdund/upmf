package models

import "net"

const (
	NET_TYPE_AN   uint8 = 0 //connect to RAN nodes
	NET_TYPE_TRAN uint8 = 1 //between two UPFs
	NET_TYPE_DNN  uint8 = 2 //UPF to DN

	SBI_DEFAULT_IP = "0.0.0.0"
)

type InfAddr interface {
	GetIpAddr() net.IP
}

type IpAddr net.IP

type DnnInfo struct {
	IpAddr    `json:"ipaddr"`
	Allocator *IpAllocator `json:"allocator"`
}

func (addr IpAddr) GetIpAddr() net.IP {
	return net.IP(addr)
}

// "infs"
type NetInf struct {
	Id      string      //unique id in a  topo, compose of the node id, the network, and the interface index
	Netname string      //network that this face connects to
	Nettype uint8       //type of network
	Addr    InfAddr     //ipv4 or ipv6
}

func (inf *NetInf) IsAn() bool {
	return inf.Nettype == NET_TYPE_AN
}

func (inf *NetInf) IsDnn() bool {
	return inf.Nettype == NET_TYPE_DNN
}
