package upf

import (
	"net/http"
	"upmf/internal/context"
	modelcontext "upmf/internal/context"
	"upmf/internal/upftopo"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func UpfDeregister(nf *modelcontext.UPMF) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var upfNodeConfig modelcontext.NodeConfig
		if err := ctx.BindJSON(&upfNodeConfig); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for _, snssai := range nf.Config.Slices {
			upftopo.RemoveNode(&upfNodeConfig, nf.TopoMaps[snssai])
			nf.TopoMaps[snssai].Links = upftopo.RemoveLink(nf.TopoMaps[snssai].Links, upfNodeConfig.Id)
		}

		ctx.JSON(http.StatusOK, gin.H{"Status": "DEREGISTERED"})
		logrus.Infoln(upfNodeConfig.Id, "is Deregisterd")

		// for _, topo := range nf.TopoMaps {
		// 	PrintLink(topo.Links)
		// }

		return
	}
}

func PrintLink(Links []context.Link) {
	for _, link := range Links {
		logrus.Infoln(link.Inf1.Id, "-", link.Inf2.Id)
	}
}
