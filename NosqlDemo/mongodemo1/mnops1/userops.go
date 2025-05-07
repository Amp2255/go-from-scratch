package mnops1

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectToDb(ctx context.Context) {
	url := "mongodb://172.31.7.157:27017/"
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println("Mongo DB not connected : ", err)
	}
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("golangdb").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

}

func CloseConnection(ctx context.Context) {
	client.Disconnect(ctx)
	fmt.Println("DB connection closed")
}

func InsertNewUser(ctx context.Context) {

	coll := client.Database("golangdb").Collection("user")

	_, err := coll.InsertOne(ctx, bson.M{"name": "John Doe", "age": 22, "gender": "Male"})
	if err != nil {
		fmt.Println("Error occured while insertion")
	} else {
		fmt.Println("Data saved")
	}

}
