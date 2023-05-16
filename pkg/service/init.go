package service

import (
	"io"
	"os"
	"upmf/internal/context"
	"upmf/internal/sbi/smf"
	"upmf/internal/sbi/upf"
	"upmf/internal/upftopo"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "service"})
}

func New(config *context.UpmfConfig) (nf *context.UPMF, err error) {
	// nf.Config = config
	nf = &context.UPMF{
		Config: config,
		// Nodes:   make(map[string]*upman.UpNode),
		ListMap: make(map[string][]context.QueryMap),
		UpfTopo: &context.UpfTopo{
			Nets:       make(map[string]uint8),
			Nodes:      make(map[string]*context.TopoNode),
			Sbiid2node: make(map[string]*context.TopoNode),
			Links:      []context.Link{},
			Heartbeat:  0,
		},
	}
	// topo := context.UpfTopo {

	// }
	return
}

func Start(nf *context.UPMF) {
	var config context.TopoConfig
	upftopo.ParseNets(nf.UpfTopo, &config)
	upftopo.ParseLinks(nf.UpfTopo, &config)
	handleSbi(nf, &config)
}

func handleSbi(nf *context.UPMF, config *context.TopoConfig) {
	gin.DisableConsoleColor()
	// Logging to a file.
	sbilog, _ := os.Create("config/sbi.log")
	gin.DefaultWriter = io.MultiWriter(sbilog)

	router := gin.Default()

	routerUpf := router.Group("/upf")
	{
		routerUpf.PUT("/register", upf.UpfRegister(nf, config))
		routerUpf.DELETE("/register", upf.UpfDeregister(nf))
		// routerUpf.PUT("/register", upf.UpfUpdate)
	}
	routerSmf := router.Group("/smf")
	{
		routerSmf.GET("/query", smf.GetQuery(nf))
	}
	router.Run(":8081")
	// router.RunTLS(":8081", "./config/TLS/server.pem", "./config/TLS/server.key")
}

func Stop(nf *context.UPMF) {
	os.Exit(0)
}
