package upf

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

var upfnode_collection *mongo.Collection

// var RegisterNode = make(chan context.TopoNode)
// var DeregisterNode = make(chan context.TopoNode)
var log *logrus.Entry

func init() {
	// Set client options
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// // Connect to MongoDB
	// client, err := mongo.Connect(context.Background(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // Check the connection
	// err = client.Ping(context.Background(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// db := client.Database("free5gc")
	// upfnode_collection = db.Collection("upfnode")

	log = logrus.WithFields(logrus.Fields{"mod": "upf"})
}

// func UpfHandler(nf *context.UPMF) {

// 	gin.DisableConsoleColor()
// 	// Logging to a file.
// 	ginupf, _ := os.Create("config/upf.log")
// 	gin.DefaultWriter = io.MultiWriter(ginupf)

// 	upf_route := gin.New()
// 	upf_route.PUT("/handleupf", UpfRegister(nf))
// 	upf_route.DELETE("/handleupf", UpfDeregister(nf))

// 	upf_route.Run(":8081")
// }
