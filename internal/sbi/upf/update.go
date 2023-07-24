package upf

import (
	"net/http"
	"upmf/internal/context"
	"upmf/internal/upftopo"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func UpfUpdate(nf *context.UPMF) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var upfNodeConfig context.NodeConfig
		if err := ctx.BindJSON(&upfNodeConfig); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Cause": err.Error()})
			return
		}

		if upfNodeConfig.Msg == "Heart-Beat" {
			ctx.JSON(http.StatusOK, gin.H{"Status": "Heart-Beat Timer"})
		}

		for _, slice := range upfNodeConfig.Slices {
			if snssai, ok := nf.Config.Slices[slice]; ok {
				upftopo.RemoveNode(&upfNodeConfig, nf.TopoMaps[snssai])
				nf.TopoMaps[snssai].Links = upftopo.RemoveLink(nf.TopoMaps[snssai].Links, upfNodeConfig.Id)
			}
		}

		upftopo.ParseNode(&upfNodeConfig, nf)

		ctx.JSON(http.StatusOK, gin.H{"Status": "UPDATED"})
		logrus.Infoln(upfNodeConfig.Id, "is Updated")
		return
	}
}
