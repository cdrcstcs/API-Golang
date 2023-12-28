package db

import (
	"context"
	"fmt"

	"github.com/tryvium-travels/memongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"inititaryplanner/common/config"
	"inititaryplanner/constant"
)

type MainMongoDB mongo.Database

func GetMainMongoDatabase() *MainMongoDB {
	return (*MainMongoDB)(GetMongoClient().Database(constant.MainMongoDB))
}

func GetMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(config.GlobalConfig.MongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	return client
}

// GetMemoMongo get the db in memory for unit testing
func GetMemoMongo(dbName string) *mongo.Database {
	mongoTestServer, err := memongo.StartWithOptions(&memongo.Options{
		CachePath:        "/tmp/",
		MongoVersion:     "4.2.1",
		ShouldUseReplica: true,
		DownloadURL:      "https://fastdl.mongodb.org/osx/mongodb-macos-x86_64-4.4.6.tgz",
	})
	if err != nil {
		panic(err)
	}

	uri := mongoTestServer.URI()
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("error init mem db %v", err) // TODO use log package here
		panic(err)
	}
	return client.Database(dbName)
}
