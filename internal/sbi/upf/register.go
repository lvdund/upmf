package upf

import (
	"net/http"
	"upmf/internal/context"
	"upmf/internal/upftopo"

	"github.com/gin-gonic/gin"
)

func UpfRegister(nf *context.UPMF) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var upfNodeConfig context.NodeConfig
		if err := ctx.BindJSON(&upfNodeConfig); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Cause": err.Error()})
			return
		}
		// filter := bson.M{"id": upfNodeConfig.Id}
		// if count, _ := upfnode_collection.CountDocuments(context.Background(), filter); count > 0 {
		// 	ctx.JSON(http.StatusOK, gin.H{"Status": "ALREADY REGISTERED"})
		// 	return
		// }
		// if _, err := upfnode_collection.InsertOne(context.Background(), upfNodeConfig); err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"Cause": err.Error()})
		// 	return
		// }
		ctx.JSON(http.StatusOK, gin.H{"Status": "REGISTERED"})
		log.Infoln(upfNodeConfig.Id, "is Registerd")

		// RegisterNode <- *ParseNode(&upfNodeConfig, nf)
		_ = upftopo.ParseNode(&upfNodeConfig, nf)
	}
}
