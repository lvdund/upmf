package upftopo

import (
	"net"
	"upmf/internal/context"
	"upmf/models"
)

func GenLink(nodeA *context.TopoNode, netInf *context.NetInf, topo *context.UpfTopo, snssai models.Snssai) {
	if netInf.Nettype != context.NET_TYPE_TRAN {
		return
	}
	for _, nodeB := range topo.Nodes {
		if nodeA.Id == nodeB.Id {
			continue
		}

		for _, infs := range nodeB.Infs {
			for _, inf := range infs {
				if inf.Nettype != context.NET_TYPE_TRAN {
					continue
				} else if !checkSameGateWay(netInf.Addr.GetIpAddr(), inf.Addr.GetIpAddr()) {
					continue
				}
				inf.Remotes = append(inf.Remotes, nodeA)
				netInf.Remotes = append(netInf.Remotes, nodeB)
				if checkSameSlice(nodeA.Slices, snssai) == true {
					topo.Links = append(topo.Links, context.Link{
						Inf1: netInf,
						Inf2: &inf,
						W:    1,
					})
					log.Infoln("Create link:", netInf.Id, "--", inf.Id, "in slice:", snssai)
				}
			}
		}
	}
}

func RemoveLink(Links []context.Link, id string, snssai models.Snssai) (resultLinks []context.Link) {
	for _, link := range Links {
		if link.Inf1.Local.Id == id || link.Inf2.Local.Id == id {
			log.Infoln("Deleted:", link.Inf1.Id, "-", link.Inf2.Id, "in slice:", snssai)
		} else {
			resultLinks = append(resultLinks, link)
		}
	}
	return
}

func checkSameGateWay(ip1 net.IP, ip2 net.IP) bool {
	mask1 := ip1.DefaultMask()
	mask2 := ip2.DefaultMask()
	return mask1.String() == mask2.String()
}

func checkSameSlice(slice1 []models.Snssai, snssai2 models.Snssai) bool {
	for _, snssai1 := range slice1 {
		if snssai1 == snssai2 {
			return true
		}
	}
	return false
}
