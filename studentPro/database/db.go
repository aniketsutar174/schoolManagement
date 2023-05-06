package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ConnectionString = "mongodb://localhost:27017"
const DbName = "studentData"
const ColName = "studentCol"

var StudCollection *mongo.Collection

func init() {
	//client option
	clientOption := options.Client().ApplyURI(ConnectionString)

	//connect mongodb
	client, _ := mongo.Connect(context.TODO(), clientOption)
	fmt.Println("MongoDB connection success")
	StudCollection = client.Database(DbName).Collection(ColName)
	fmt.Println("MongoDB connection successfull")

}
