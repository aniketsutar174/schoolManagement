package service

import (
	"context"
	"log"

	"github.com/aniketsutar174/schoolManagement/database"
	"github.com/aniketsutar174/schoolManagement/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AllDataStud(studData []model.StudDatabase) ([]model.StudDatabase, error) {
	curser, _ := database.StudCollection.Find(context.Background(), bson.D{{}})
	for curser.Next(context.Background()) {
		var stud model.StudDatabase
		err := curser.Decode(&stud)
		if err != nil {
			log.Fatal(err)
		}
		studData = append(studData, stud)
	}
	defer curser.Close(context.Background())
	return studData, nil
}

func AllDataByIdStud(id string) (primitive.M, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	var result bson.M
	err := database.StudCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}
func UpdateStudDataByIdStud(id string, studData model.StudDatabase) (*mongo.UpdateResult, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	updateQuery := bson.M{
		"marks.maths":   studData.Marks.Maths,
		"marks.science": studData.Marks.Science,
	}
	update := bson.M{
		"$set": updateQuery,
	}
	result, _ := database.StudCollection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	return result, nil
}
func InsertDataStud(studData model.StudDatabase) (*mongo.InsertOneResult, error) {
	studData.Id = primitive.NewObjectID()
	result, err := database.StudCollection.InsertOne(context.Background(), studData)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func DeleteDataById(id string) (*mongo.DeleteResult, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := database.StudCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return nil, err
	}
	log.Println("deleted Id:", id)
	return result, nil
}
