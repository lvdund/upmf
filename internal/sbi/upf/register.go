package upf

import (
	"net/http"
	"upmf/internal/context"
	"upmf/internal/upftopo"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func UpfRegister(nf *context.UPMF) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var upfNodeConfig context.NodeConfig
		if err := ctx.BindJSON(&upfNodeConfig); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Cause": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"Status": "REGISTERED"})
		logrus.Infoln(upfNodeConfig.Id, "is Registerd")

		upftopo.ParseNode(&upfNodeConfig, nf)
	}
}
