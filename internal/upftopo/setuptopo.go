package upftopo

// func SetupUpfTopo(topo *context.UpfTopo, node *context.TopoNode) {
// 	SetupHeartbeat(topo, node)
// 	SetupNets(topo, node)
// 	// SetupNodes(topo, node)
// 	// SetupPfcpid2node(topo, node)
// 	// SetupLinks(topo)
// }

// func SetupHeartbeat(topo *context.UpfTopo, node *context.TopoNode) {
// 	if topo.Heartbeat == 0 {
// 		topo.Heartbeat = node.Heartbeat
// 	}
// 	if topo.Heartbeat > node.Heartbeat {
// 		topo.Heartbeat = node.Heartbeat
// 	}
// }

// func SetupNets(topo *context.UpfTopo, node *context.TopoNode) {
// 	for _, infs := range node.Infs {
// 		for _, inf := range infs {
// 			topo.Nets[inf.Netname] = inf.Nettype
// 		}
// 	}
// }

// func SetupNodes(topo *context.UpfTopo, node *context.TopoNode) {
// 	topo.Nodes[node.Id] = node
// }

// func SetupPfcpid2node(topo *context.UpfTopo, node *context.TopoNode) {
// 	if node.HasSbiIp() {
// 		topo.Sbiid2node[node.Sbiinfo.NodeId()] = node
// 	}
// }

// func SetupLinks(topo *context.UpfTopo) {
// 	var (
// 		// a, b                 *topoNode
// 		inf1, inf2 *context.NetInf
// 		// ok                   bool
// 		// ntype                uint8
// 		// aindex, bindex       int
// 		w uint16
// 		// linkname1, linkname2 string
// 		// err                  error
// 	)
// 	topo.Links = append(topo.Links, context.Link{
// 		Inf1: inf1,
// 		Inf2: inf2,
// 		W:    w,
// 	})
// }

// // Get a node's network interfaces to Access networks
// func GetNodeAnFaces(topo *context.UpfTopo, node *context.TopoNode, nets []string) (foundinfs []context.NetInf) {
// 	for network, infs := range node.Infs {
// 		if ntype, ok := topo.Nets[network]; ok && ntype == context.NET_TYPE_AN {
// 			for _, netname := range nets {
// 				if strings.Compare(netname, network) == 0 {
// 					foundinfs = append(foundinfs, infs...)
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return
// }

// // Get a node's network interfaces to Dnn
// func GetNodeDnnFaces(topo *context.UpfTopo, node *context.TopoNode, dnn string) (foundinfs []context.NetInf) {
// 	for network, infs := range node.Infs {
// 		if ntype, ok := topo.Nets[network]; ok && ntype == context.NET_TYPE_DNN && strings.Compare(network, dnn) == 0 {
// 			foundinfs = infs
// 			break
// 		}
// 	}
// 	return
// }
