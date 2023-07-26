package upftopo

import (
	"fmt"
	"net"
	"upmf/internal/context"
	"upmf/internal/util/ipalloc"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"upmf": "topo"})
}

func ParseNode(upfNodeConfig *context.NodeConfig, nf *context.UPMF) {
	node := context.NewNode(upfNodeConfig.Id, upfNodeConfig.Sbi.Heartbeat, upfNodeConfig.Sbi, true)
	for _, slice := range upfNodeConfig.Slices {
		if snssai, ok := nf.Config.Slices[slice]; ok {
			node.Slices = append(node.Slices, snssai)
		}
	}

	var infs []context.NetInf
	for netname, inf := range upfNodeConfig.Infs {
		if nettype, ok := nf.Config.Nets[netname]; ok {
			for _, snssai := range node.Slices {
				nf.TopoMaps[snssai].Nets[netname] = nettype
			}
			infs = []context.NetInf{}
			if nettype == context.NET_TYPE_DNN {
				dnninfolist := ParseDnnInfoList(inf)
				for i, addr := range dnninfolist {
					addInf := context.NetInf{
						Id:      fmt.Sprintf("%s:%s:%d", node.Id, netname, i),
						Netname: netname,
						Nettype: nettype,
						Addr:    addr,
						Local:   node,
					}
					infs = append(infs, addInf)
				}
			} else {
				ipaddrlist := ParseIpAddrList(inf)
				for i, addr := range ipaddrlist {
					addInf := context.NetInf{
						Id:      fmt.Sprintf("%s:%s:%d", node.Id, netname, i),
						Netname: netname,
						Nettype: nettype,
						Addr:    addr,
						Local:   node,
					}
					infs = append(infs, addInf)
					for _, snssai := range node.Slices {
						GenLink(node, &addInf, nf.TopoMaps[snssai], snssai)
					}
				}
			}
			node.Infs[netname] = infs
		} else {
			// log
		}
	}
	for _, snssai := range node.Slices {
		nf.TopoMaps[snssai].Nodes[node.Id] = node
		log.Infoln("Add UPF", node.Id, "in slice:", snssai)
		if node.HasSbiIp() {
			nf.TopoMaps[snssai].Sbiid2node[node.Sbiinfo.NodeId()] = node
		} else {
			log.Infoln("No sbi to", node.Id)
		}
	}

	return
}

func RemoveNode(upfId string, topo *context.UpfTopo) {
	delete(topo.Nodes, upfId)

	for sbi, node := range topo.Sbiid2node {
		if node.Id == upfId {
			delete(topo.Sbiid2node, sbi)
		}
	}
}

func ParseIpAddrList(infs []context.NetInfConfig) (addrlist []context.IpAddr) {
	for _, info := range infs {
		if ip := net.ParseIP(info.Addr); ip != nil {
			addrlist = append(addrlist, context.IpAddr(ip))
		} else {
			log.Warnf("parse IP fails '%s'", info.Addr)
		}
	}
	return
}

func ParseDnnInfoList(infs []context.NetInfConfig) (addrlist []*context.DnnInfo) {
	for _, info := range infs {
		if info.DnnInfo != nil {
			if ip := net.ParseIP(info.Addr); ip != nil {
				if _, ipnet, err := net.ParseCIDR(info.DnnInfo.Cidr); err == nil && ipnet != nil {
					addrlist = append(addrlist, &context.DnnInfo{
						IpAddr:    context.IpAddr(ip),
						Allocator: ipalloc.New(ipnet),
					})
					// log.Infoln("info.Addr:", info.Addr)
					// log.Infoln("addrlist:", info.DnnInfo.Cidr, "to ->", ip, "-", ipnet.IP, net.IP(ipnet.Mask))
				} else if err != nil {
					log.Warnf("Parse CIDR %s returns error: %s", info.DnnInfo.Cidr, err.Error())
				}
			} else {
				log.Warnf("parse IP fails '%s'", info.Addr)
			}
		} else {
			log.Warnf("Dnn at %s has empty cidr", info.Addr)
		}
	}
	// for _, value := range addrlist {
	// 	log.Infoln("Check parseDnnInfoList:", value.IpAddr(), value.allocator.GetCidr())
	// }
	return
}
