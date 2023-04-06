package db

// ToDO: Create a database connection and get the collection

import (
	"context"
	"fmt"
	"gin/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Hello struct {
	ID primitive.ObjectID `bson:"_id"`
	message string
}



func GetHelloString() string {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(config.GetConfig().ConnectionString))
	// if err != nil { return err }

	client.Database(config.GetConfig().DatabaseName).CreateCollection(ctx, "hello", nil)

	collection := client.Database(config.GetConfig().DatabaseName).Collection("hello")
	
	filter := bson.D{{Key: "message", Value: "Hello World!"}}

	var result Hello

	collection.FindOne(ctx, filter).Decode(&result)

	// if err != nil {
	// 	doc := bson.D{{Key: "message", Value: "Hello World!"}}
	// 	collection.InsertOne(context.TODO(), doc)
	// 	panic(err)
	// }

	fmt.Print("result: " + result.message)

	return result.message;
}

// Create golang mongo client and get collection field