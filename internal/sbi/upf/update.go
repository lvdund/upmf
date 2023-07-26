package upf

import (
	"net/http"
	"upmf/internal/context"
	"upmf/internal/upftopo"

	"github.com/gin-gonic/gin"
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

		upfId := ctx.Param("nameID")

		for _, slice := range upfNodeConfig.Slices {
			if snssai, ok := nf.Config.Slices[slice]; ok {
				upftopo.RemoveNode(upfId, nf.TopoMaps[snssai])
				for i, link := range nf.TopoMaps[snssai].Links {
					if link.Inf1.Local.Id == upfId || link.Inf2.Local.Id == upfId {
						nf.TopoMaps[snssai].Links = append(nf.TopoMaps[snssai].Links[:i], nf.TopoMaps[snssai].Links[i+1:]...)
					}
				}
			}
		}

		upftopo.ParseNode(&upfNodeConfig, nf)

		ctx.JSON(http.StatusOK, gin.H{"Status": "UPDATED"})
		log.Infoln(upfNodeConfig.Id, "is Updated")
		return
	}
}
