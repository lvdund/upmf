package service

import (
	"io"
	"os"
	"upmf/internal/context"
	"upmf/internal/sbi/smf"
	"upmf/internal/sbi/upf"
	"upmf/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "service"})
}

func New(config *context.UpmfConfig) (nf *context.UPMF, err error) {
	nf = &context.UPMF{
		Config:   config,
		TopoMaps: make(map[models.Snssai]*context.UpfTopo),
	}
	for _, snssai := range config.Slices {
		nf.TopoMaps[snssai] = &context.UpfTopo{
			Nets:       make(map[string]uint8),
			Nodes:      make(map[string]*context.TopoNode),
			Sbiid2node: make(map[string]*context.TopoNode),
			Links:      []context.Link{},
			Heartbeat:  3,
		}
	}
	return
}

func Start(nf *context.UPMF) {
	handleSbi(nf)
}

func handleSbi(nf *context.UPMF) {
	gin.DisableConsoleColor()
	// Logging to a file.
	sbilog, _ := os.Create("config/sbi.log")
	gin.DefaultWriter = io.MultiWriter(sbilog)

	router := gin.Default()

	routerUpf := router.Group("/upf")
	{
		routerUpf.PUT("/register", upf.UpfRegister(nf))
		routerUpf.PATCH("/register", upf.UpfUpdate(nf))
		routerUpf.DELETE("/register", upf.UpfDeregister(nf))
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
