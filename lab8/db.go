package db

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB() *mongo.Client {
	clientOptions := options.Client().ApplyURI( uri: "mongodb://172.17.0.2:27017" )		//connect to mongoDB
	client, err := mongo.connect(context.TODO(), clientOptions)
	if err != nil {
		log.fatal(err)
	}

	//check the connection
	err= client.Ping(context.TODO(), rp: nil)
	if err != nil {
		log.fatal(err)
	}
	fmt.Println(a...: "connect to MongoDB!")
	return client
}