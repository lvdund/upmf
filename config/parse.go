package config

// func parseIpAddrList(infs []models.NetInfConfig) (addrlist []models.InfAddr) {
// 	for _, info := range infs {
// 		if ip := net.ParseIP(info.Addr); ip != nil {
// 			addrlist = append(addrlist, models.IpAddr(ip))
// 		} else {
// 			logrus.Warnf("parse IP fails '%s'", info.Addr)
// 		}
// 	}
// 	return
// }

// func parseDnnInfoList(infs []models.NetInfConfig) (addrlist []*models.DnnInfo) {
// 	for _, info := range infs {
// 		if info.DnnInfo != nil {
// 			if ip := net.ParseIP(info.Addr); ip != nil {
// 				if _, ipnet, err := net.ParseCIDR(info.DnnInfo.Cidr); err == nil && ipnet != nil {
// 					addrlist = append(addrlist, &models.DnnInfo{
// 						IpAddr:    models.IpAddr(ip),
// 						Allocator: models.New(ipnet),
// 					})
// 				} else if err != nil {
// 					logrus.Warnf("Parse CIDR %s returns error: %s", info.DnnInfo.Cidr, err.Error())
// 				}
// 			} else {
// 				logrus.Warnf("parse IP fails '%s'", info.Addr)
// 			}
// 		} else {
// 			logrus.Warnf("Dnn at %s has empty cidr", info.Addr)
// 		}
// 	}

// 	return
// }
