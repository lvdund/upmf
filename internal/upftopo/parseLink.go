package upftopo

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"upmf/internal/context"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "config topo"})
}

func ParseNets(topo *context.UpfTopo, config *context.TopoConfig) {
	var err error
	var buf []byte
	if buf, err = ioutil.ReadFile("config/nets_links.json"); err != nil {
		// log
		return
	}
	if err = json.Unmarshal(buf, &config); err != nil {
		// log
		return
	}

	for _, name := range config.Networks.Access { // name: ["an1", "an2"]
		topo.Nets[name] = context.NET_TYPE_AN
	}

	for _, name := range config.Networks.Transport { // ["tran"]
		topo.Nets[name] = context.NET_TYPE_TRAN
	}

	for _, name := range config.Networks.Dnn { // ["e1", "e2", "internet"]
		topo.Nets[name] = context.NET_TYPE_DNN
	}
}

// func ParseLinks(topo *context.UpfTopo, config *context.TopoConfig) {
// 	var (
// 		a, b                 *context.TopoNode
// 		inf1, inf2           *context.NetInf
// 		ok                   bool
// 		ntype                uint8
// 		aindex, bindex       int
// 		w                    uint16
// 		linkname1, linkname2 string
// 		err                  error
// 	)
// 	// Links
// 	existedlinks := make(map[string]uint16) //mark parsed links

// 	//loop through all networks to parse links
// 	for netname, linklist := range config.Links { // "transport": ["tran"]
// 		// netname = "tran"
// 		//check if network is defined and it is a transport network
// 		if ntype, ok = topo.Nets[netname]; !ok || ntype != context.NET_TYPE_TRAN {
// 			// ntype = 1 -- tran
// 			log.Warnf("'%s' either not exist or not a transport network", netname)
// 			continue
// 		}
// 		//parse all links in this network
// 		for _, linkconfig := range linklist {
// 			if a, ok = topo.Nodes[linkconfig.A.Node]; !ok {
// 				log.Warnf("'%s' not exist", linkconfig.A.Node)
// 				continue
// 			}
// 			if b, ok = topo.Nodes[linkconfig.B.Node]; !ok {
// 				log.Warnf("'%s' not exist", linkconfig.B.Node)
// 				continue
// 			}
// 			//make sure a and b are different nodes
// 			if strings.Compare(a.Id, b.Id) == 0 {
// 				log.Warnf("link with duplicated endpoint '%s'", a.Id)
// 				continue
// 			}
// 			//get interface indexes
// 			aindex = 0
// 			if linkconfig.A.Index != nil {
// 				aindex = *linkconfig.A.Index
// 			}
// 			bindex = 0
// 			if linkconfig.B.Index != nil {
// 				bindex = *linkconfig.B.Index
// 			}
// 			//get link weight
// 			w = 1 //default link weight
// 			if linkconfig.W != nil {
// 				w = *linkconfig.W
// 			}

// 			if inf1, err = a.GetInf(netname, aindex); err != nil {
// 				// log.Warnf("get face from node '%s' fails: %s", a.Id, err.Error())
// 				continue
// 			}
// 			// cc, _ := a.getInf(netname, aindex)
// 			// // log.Info(cc.addr.IpAddr().String())
// 			if inf2, err = b.GetInf(netname, bindex); err != nil {
// 				// log.Warnf("get face from node '%s' fails: %s", b.Id, err.Error())
// 				continue
// 			}
// 			// // log.Infoln("-----", inf2.addr.IpAddr())
// 			//check if link existed already
// 			linkname1 = fmt.Sprintf("%s:%s:%d-%s:%s:%d", a.Id, netname, aindex, b.Id, netname, bindex)
// 			linkname2 = fmt.Sprintf("%s:%s:%d-%s:%s:%d", b.Id, netname, bindex, a.Id, netname, aindex)
// 			// // log.Infoln("linkname1:", linkname1)
// 			// // log.Infoln("linkname2:", linkname2)
// 			if _, ok := existedlinks[linkname1]; ok {
// 				// log.Warnf("link '%s' existed", linkname1)
// 				continue
// 			}
// 			//mark link as existed
// 			existedlinks[linkname1] = w
// 			existedlinks[linkname2] = w
// 			log.Infof("add links: [%s] (alias[%s]) with weight=%d", linkname1, linkname2, w)
// 			inf1.Remotes = append(inf1.Remotes, b)
// 			inf2.Remotes = append(inf2.Remotes, a)
// 			topo.Links = append(topo.Links, context.Link{
// 				Inf1: inf1,
// 				Inf2: inf2,
// 				W:    w,
// 			})
// 		}
// 	}
// }

func GenLink(nodeA *context.TopoNode, netInf *context.NetInf, topo *context.UpfTopo) {
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
				topo.Links = append(topo.Links, context.Link{
					Inf1: netInf,
					Inf2: &inf,
					W:    1,
				})
				log.Infoln(netInf.Id, "-----", inf.Id)
			}
		}
	}
}

func checkSameGateWay(ip1 net.IP, ip2 net.IP) bool {
	mask1 := ip1.DefaultMask()
	mask2 := ip2.DefaultMask()
	return mask1.String() == mask2.String()
}
