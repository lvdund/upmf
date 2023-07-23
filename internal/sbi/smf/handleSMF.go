package smf

import (
	"fmt"
	"net/http"
	"upmf/internal/context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetQuery(nf *context.UPMF) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// var query context.Query
		var query context.PathQuery
		if err := ctx.BindJSON(&query); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Cause": err.Error()})
			return
		}

		logrus.Infoln("Query msg:", query)

		var queryPath context.DataPath

		if _, ok := nf.TopoMaps[query.Snssai]; ok {
			// if true {
			queryPath = FindPath(nf.TopoMaps[query.Snssai], &query)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"Cause": fmt.Sprintf("Cannot access to slice %s", query.Snssai.Sst)})
			return
		}
		logrus.Infoln("Response Path:", queryPath.Path)
		if queryPath.Path == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Cause": "Bad Request"})
			return
		}
		ctx.JSON(http.StatusOK, queryPath)
		return
	}
}
