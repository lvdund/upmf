package smf

import (
	"encoding/json"
	"fmt"
	"net/http"
	"upmf/internal/context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// func SmfHandler(nf *context.UPMF) {

// 	gin.DisableConsoleColor()
// 	// Logging to a file.
// 	ginsmf, _ := os.Create("config/smf.log")
// 	gin.DefaultWriter = io.MultiWriter(ginsmf)

// 	smf_route := gin.New()
// 	smf_route.PUT("/path", GetQuery(nf))
// 	smf_route.Run(":8082")
// }

func GetQuery(nf *context.UPMF) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var query context.Query
		if err := ctx.BindJSON(&query); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Cause": err.Error()})
			return
		}

		var quermap context.QueryMap

		for _, path := range nf.ListQueryMap[fmt.Sprintf("%s:%s", query.SmfId, query.UeId)] {
			path1, _ := json.Marshal(path.Query)
			path2, _ := json.Marshal(query.Query)
			if string(path1) == string(path2) {
				logrus.Infoln("Response Path:", path.Paths.Path)
				ctx.JSON(http.StatusOK, path.Paths)
				return
			}
		}
		var path context.DataPath
		if _, ok := nf.ListLinks[query.Query.Snssai]; ok {
			path = FindPath(nf.UpfTopo, nf.ListLinks[query.Query.Snssai], &query.Query)
			quermap.Paths = path
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"Cause": "loss snssai info"})
			return
		}
		logrus.Infoln("Response Path:", quermap.Paths.Path)
		quermap.Query = query.Query
		quermap.UeId = query.UeId

		nf.ListQueryMap[fmt.Sprintf("%s:%s", query.SmfId, query.UeId)] = append(nf.ListQueryMap[fmt.Sprintf("%s:%s", query.SmfId, query.UeId)], quermap)

		ctx.JSON(http.StatusOK, path)
	}
}
