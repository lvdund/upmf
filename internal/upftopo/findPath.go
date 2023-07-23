package upftopo

import (
	"fmt"
	"math/rand"
	"net"
	"upmf/internal/context"
	"upmf/internal/util/dijkstra"

	"github.com/sirupsen/logrus"
)

func FindPath(topo *context.UpfTopo, query *context.PathQuery) (datapath context.DataPath) {

	//find all anchors and source nodes for searching(at the same time)
	dnnfaces := []context.NetInf{} //Net interfaces to Dnn
	srcfaces := []context.NetInf{} //nodes for start searching

	for _, node := range topo.Nodes {
		// if node.IsActive() && node.Serve(query.Snssai) {
		if node.IsActive() {
			if infs := topo.GetNodeDnnFaces(node, query.Dnn); len(infs) > 0 {
				dnnfaces = append(dnnfaces, infs...)
			}

			if infs := topo.GetNodeAnFaces(node, query.Nets); len(infs) > 0 { //a starting node
				srcfaces = append(srcfaces, infs...)
			}
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
		logrus.Errorf("can't select an anchor to allocate Ue's IP")
		return
	}
	logrus.Infof("UE's IP = %s(%d) on Dnn=%s", ip.String(), len(ip), dnnface.Netname)

	//build a graph of active links then find the shortest paths from source to destination
	edges := []dijkstra.EdgeInfo{} //edges to build the grap

	//a structure to keep the endpoint's ip addresses of a link
	type edgesig struct {
		ip1 net.IP
		ip2 net.IP
	}

	ipmap := make(map[string]edgesig) //map edge name to a tuple of its endpoint's ip addresses

	for _, l := range topo.Links {
		if l.IsActive(query.Snssai) { //only pick active links
			edges = append(edges, dijkstra.EdgeInfo{
				A: l.Inf1.Local.Id,
				B: l.Inf2.Local.Id,
				W: int64(l.W),
			})
			logrus.Infof("add link %s-%s", l.Inf1.Local.Id, l.Inf2.Local.Id)
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
		logrus.Infof("Search path from %s to %s", srcface.Local.Id, dnnface.Local.Id)
		if _, paths := graph.ShortestPath(srcface.Local.Id, dnnface.Local.Id); len(paths) > 0 {
			path := paths[0] //pick the first path
			for _, pth := range paths {
				if len(pth) < len(path) {
					path = pth
				}
			}

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
			break
		}
	}
	return
}
