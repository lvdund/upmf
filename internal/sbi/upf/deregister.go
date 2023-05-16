package upf

import (
	"context"
	"fmt"
	"net"
	"net/http"
	modelcontext "upmf/internal/context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpfDeregister(nf *modelcontext.UPMF) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var upfNodeConfig modelcontext.NodeConfig
		if err := ctx.BindJSON(&upfNodeConfig); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		filter := bson.M{"id": upfNodeConfig.Id}
		if count, _ := upfnode_collection.CountDocuments(context.Background(), filter); count == 0 {
			ctx.JSON(http.StatusOK, gin.H{"Status": "ALREADY DEREGISTERED"})
			return
		}
		_, err := upfnode_collection.DeleteOne(context.Background(), filter)
		if err != nil {
			// log.Fatal(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		RemoveNode(&upfNodeConfig, nf)
		nf.UpfTopo.Links = RemoveLink(nf.UpfTopo.Links, upfNodeConfig.Id)
		ctx.JSON(http.StatusOK, gin.H{"Status": "DEREGISTERED"})
		log.Infoln(upfNodeConfig.Id, "is Deregisterd")
		// DeregisterNode <- upfNode
		return
	}
}

func RemoveNode(upfNodeConfig *modelcontext.NodeConfig, nf *modelcontext.UPMF) {
	delete(nf.UpfTopo.Nodes, upfNodeConfig.Id)

	var ip net.IP
	var port int
	if upfNodeConfig.Sbi != nil {
		ip = net.ParseIP(upfNodeConfig.Sbi.Ip)
		port = upfNodeConfig.Sbi.Port
	}
	delete(nf.UpfTopo.Sbiid2node, fmt.Sprintf("%s:%d", ip.String(), port))
}

func RemoveLink(Links []modelcontext.Link, id string) []modelcontext.Link {
	for i, link := range Links {
		if link.Inf1.Local.Id == id || link.Inf2.Local.Id == id {
			Links = append(Links[:i], Links[i+1:]...)
		}
	}
	return Links
}
