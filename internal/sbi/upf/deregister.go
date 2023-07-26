package upf

import (
	"net/http"
	"upmf/internal/context"
	modelcontext "upmf/internal/context"
	"upmf/internal/upftopo"

	"github.com/gin-gonic/gin"
)

func UpfDeregister(nf *modelcontext.UPMF) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		upfId := ctx.Param("nameID")

		for _, snssai := range nf.Config.Slices {
			upftopo.RemoveNode(upfId, nf.TopoMaps[snssai])
			nf.TopoMaps[snssai].Links = upftopo.RemoveLink(nf.TopoMaps[snssai].Links, upfId, snssai)
		}
		log.Infoln("Deleted UPF", upfId)

		ctx.JSON(http.StatusOK, gin.H{"Status": "DEREGISTERED"})
		log.Infoln(upfId, "is Deregisterd")

		return
	}
}

func PrintLink(Links []context.Link) {
	for _, link := range Links {
		log.Infoln(link.Inf1.Id, "-", link.Inf2.Id)
	}
}
