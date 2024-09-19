package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongo() *Mongodb {
	return &Mongodb{
		uri: os.Getenv("MONGO_DB_URL"),
	}
}

func (m *Mongodb) InitMongo() {
	// fetch todo context
	ctx := context.TODO()
	mongoconn := options.Client().ApplyURI(m.uri)
	// connect with the mongo server
	mongodb, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("Cannot initialize mongo")
	}
	// ping the server to check if the connection is successful or not
	err = mongodb.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error connecting with mongo")
	}
	fmt.Println("Mongo Connection Established")
	// select the database of the application
	m.DB = mongodb.Database("event-manager")
	m.ctx = &ctx

}
