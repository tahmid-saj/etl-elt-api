package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBClient *mongo.Client
var err error

func InitMongoDB() {
  // Use the SetServerAPIOptions() method to set the version of the Stable API on the client
  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := fmt.Sprintf("%v%v%v", os.Getenv("MONGODB_URL_1"), os.Getenv("MONGODB_PASSWORD"), os.Getenv("MONGODB_URL_2"))
  opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

  // Create a new client and connect to the server
  MongoDBClient, err = mongo.Connect(context.TODO(), opts)
	fmt.Println(MongoDBClient)
  if err != nil {
    panic(err)
  }

	fmt.Println("successfully initiated connection to mongodb")
}

func DisconnectMongoDB() {
	fmt.Println("closing mongodb connection")
  defer func() {
    if err := MongoDBClient.Disconnect(context.TODO()); err != nil {
      panic(err)
    }
  }()
}