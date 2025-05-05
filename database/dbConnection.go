package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DbInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

	uri := os.Getenv("MONGODB_URL")
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	println("MONGO IS RUNNING")
	return client
}

var Client = DbInstance()

func OpenCollection(collname *mongo.Client, dbname string) *mongo.Collection {
	coll := collname.Database("clustor0").Collection(dbname)

	return coll
}
