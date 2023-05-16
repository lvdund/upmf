package upf

import (
	"context"
	"net/http"
	modelcontext "upmf/internal/context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpfUpdate(ctx *gin.Context) {
	var upfNode modelcontext.NodeConfig
	if err := ctx.BindJSON(&upfNode); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Cause": err.Error()})
		return
	}
	filter := bson.M{"id": upfNode.Id}

	if count, _ := upfnode_collection.CountDocuments(context.Background(), filter); count > 0 {
		return
	}
	if _, err := upfnode_collection.ReplaceOne(context.Background(), filter, upfNode); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Cause": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Status": "UPDATED"})
	// RegisterNode <- upfNode
	return
}
