package upf

import (
	"fmt"
	"net"
	"net/http"
	"upmf/internal/context"
	"upmf/internal/upftopo"
	"upmf/internal/util/ipalloc"

	"github.com/gin-gonic/gin"
)

func UpfRegister(nf *context.UPMF, config *context.TopoConfig) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var upfNodeConfig context.NodeConfig
		if err := ctx.BindJSON(&upfNodeConfig); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Cause": err.Error()})
			return
		}
		// filter := bson.M{"id": upfNodeConfig.Id}
		// if count, _ := upfnode_collection.CountDocuments(context.Background(), filter); count > 0 {
		// 	ctx.JSON(http.StatusOK, gin.H{"Status": "ALREADY REGISTERED"})
		// 	return
		// }
		// if _, err := upfnode_collection.InsertOne(context.Background(), upfNodeConfig); err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"Cause": err.Error()})
		// 	return
		// }
		ctx.JSON(http.StatusOK, gin.H{"Status": "REGISTERED"})
		log.Infoln(upfNodeConfig.Id, "is Registerd")

		// RegisterNode <- *ParseNode(&upfNodeConfig, nf)
		_ = ParseNode(&upfNodeConfig, nf)
		// upftopo.ParseLinks(nf.UpfTopo, config)
		// return
	}
}

func ParseNode(upfNodeConfig *context.NodeConfig, nf *context.UPMF) (node *context.TopoNode) {
	node = context.NewNode(upfNodeConfig.Id, upfNodeConfig.Sbi.Heartbeat, upfNodeConfig.Sbi, true)
	for _, slice := range upfNodeConfig.Slices {
		if snssai, ok := nf.Config.Slices[slice]; ok {
			node.Slices = append(node.Slices, snssai)
		}
	}

	var infs []context.NetInf
	for netname, inf := range upfNodeConfig.Infs {
		if nettype, ok := nf.UpfTopo.Nets[netname]; ok {
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
					// log.Infof(fmt.Sprintf("%s:%s:%d", node.Id, netname, i))
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
					upftopo.GenLink(node, &addInf, nf.UpfTopo, nf.ListLinks)
					// log.Infof(fmt.Sprintf("%s:%s:%d", node.Id, netname, i))
				}
			}
			node.Infs[netname] = infs
		} else {
			// log
		}
		nf.UpfTopo.Nodes[node.Id] = node
		if node.HasSbiIp() {
			nf.UpfTopo.Sbiid2node[node.Sbiinfo.NodeId()] = node
		} else {
			log.Infoln("No sbi to", node.Id)
		}
	}

	return
}

func ParseIpAddrList(infs []context.NetInfConfig) (addrlist []context.IpAddr) {
	for _, info := range infs {
		if ip := net.ParseIP(info.Addr); ip != nil {
			addrlist = append(addrlist, context.IpAddr(ip))
		} else {
			// log.Warnf("parse IP fails '%s'", info.Addr)
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
					// log.Warnf("Parse CIDR %s returns error: %s", info.DnnInfo.Cidr, err.Error())
				}
			} else {
				// log.Warnf("parse IP fails '%s'", info.Addr)
			}
		} else {
			// log.Warnf("Dnn at %s has empty cidr", info.Addr)
		}
	}
	// for _, value := range addrlist {
	// 	log.Infoln("Check parseDnnInfoList:", value.IpAddr(), value.allocator.GetCidr())
	// }
	return
}
