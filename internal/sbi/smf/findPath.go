package smf

import (
	"fmt"
	"math/rand"
	"net"
	"upmf/internal/context"
	"upmf/internal/util/dijkstra"
)

func FindPath(topo *context.UpfTopo, listlink []context.Link, query *context.PathQuery) (datapath context.DataPath) {
	// log := logrus.WithFields(logrus.Fields{"mod": "path"})
	// log.Infoln(query)
	//find all anchors and source nodes for searching(at the same time)
	dnnfaces := []context.NetInf{} //Net interfaces to Dnn
	srcfaces := []context.NetInf{} //nodes for start searching
	var prtstr string

	for _, node := range topo.Nodes {

		// for nameInfs, _ := range node.Infs {
		// 	log.Infoln(node.Id,":", nameInfs)
		// }

		if node.IsActive() && node.Serve(query.Snssai) {
			if infs := topo.GetNodeDnnFaces(node, query.Dnn); len(infs) > 0 {
				// for _, inf := range infs {
				// 	log.Infoln(inf.Id)
				// }
				dnnfaces = append(dnnfaces, infs...)
			}

			if infs := topo.GetNodeAnFaces(node, query.Nets); len(infs) > 0 { //a starting node
				// for _, inf := range infs {
				// 	log.Infoln(inf.Id)
				// }
				srcfaces = append(srcfaces, infs...)
			}
		} else {
			// log.Infoln("Nodes have not activated:", node.Isactive, node.Serve(query.Snssai))
		}
	}

	//select an dnn face and allocate ip for UE
	var dnnface *context.NetInf //selected dnn face
	var ip net.IP
	//	1. first shuffling anchors
	for i := range dnnfaces {
		j := rand.Intn(i + 1)
		dnnfaces[i], dnnfaces[j] = dnnfaces[j], dnnfaces[i]
	}
	//	2. then pick the first one that can allocate an Ip address
	for _, face := range dnnfaces {
		dnninfo := face.Addr.(*context.DnnInfo) //must not panic
		if ip = dnninfo.Allocator.Allocate(); ip != nil {
			dnnface = &face
			break
		}
	}
	if dnnface == nil {
		// log.Errorf("can't select an anchor to allocate Ue's IP")
		return
	}
	// log.Infof("UE's IP = %s(%d) on Dnn=%s", ip.String(), len(ip), dnnface.Netname)

	//build a graph of active links then find the shortest paths from source to destination
	edges := []dijkstra.EdgeInfo{} //edges to build the grap

	//a structure to keep the endpoint's ip addresses of a link
	type edgesig struct {
		ip1 net.IP
		ip2 net.IP
	}

	ipmap := make(map[string]edgesig) //map edge name to a tuple of its endpoint's ip addresses

	for _, l := range listlink {
		if l.IsActive(query.Snssai) { //only pick active links
			edges = append(edges, dijkstra.EdgeInfo{
				A: l.Inf1.Local.Id,
				B: l.Inf2.Local.Id,
				W: int64(l.W),
			})
			// log.Infof("add link %s-%s", l.Inf1.local.id, l.Inf2.local.id)
			//keep the ip addresses of the edges for later use
			ipmap[fmt.Sprintf("%s-%s", l.Inf1.Local.Id, l.Inf2.Local.Id)] = edgesig{
				ip1: l.Inf1.Addr.GetIpAddr(),
				ip2: l.Inf2.Addr.GetIpAddr(),
			}
			ipmap[fmt.Sprintf("%s-%s", l.Inf2.Local.Id, l.Inf1.Local.Id)] = edgesig{
				ip1: l.Inf2.Addr.GetIpAddr(),
				ip2: l.Inf1.Addr.GetIpAddr(),
			}

		}
	}
	graph := dijkstra.New(edges)
	//	1. shuffle sources
	for i := range srcfaces {
		j := rand.Intn(i + 1)
		srcfaces[i], srcfaces[j] = srcfaces[j], srcfaces[i]
	}
	for _, srcface := range srcfaces {
		// log.Infof("Search path from %s to %s", srcface.local.id, dnnface.local.id)
		if _, paths := graph.ShortestPath(srcface.Local.Id, dnnface.Local.Id); len(paths) > 0 {
			// log.Infoln("Shortest path:", paths)
			path := paths[0] //pick the first path

			//build the path with ip address of the faces
			plen := len(path)
			pathnodes := make([]*context.PathNode, plen)
			for i, id := range path {
				sbiinfo := topo.Nodes[id].Sbiinfo
				pathnodes[i] = &context.PathNode{
					Id:      id,
					SbiIp:   sbiinfo.Ip,
					SbiPort: sbiinfo.Port,
				}
				// log.Infoln("node: ", id)
			}
			//set ip addresses for the An face and Dnn face of the path
			pathnodes[0].DlIp = srcface.Addr.GetIpAddr()
			pathnodes[plen-1].UlIp = dnnface.Addr.GetIpAddr()
			//set ip addresses for remaining faces on the path
			for i := 0; i < plen-1; i++ {
				info := ipmap[fmt.Sprintf("%s-%s", path[i], path[i+1])]
				pathnodes[i].UlIp = info.ip1
				pathnodes[i+1].DlIp = info.ip2
			}
			datapath = context.DataPath{
				Path: pathnodes,
				Ip:   ip,
				// Deallocator: dnnface.Addr.(*context.DnnInfo).Allocator.Release,
			}
			for _, nodepath := range pathnodes {
				prtstr += nodepath.Id + " "
			}
			break
		}
	}
	return
}
