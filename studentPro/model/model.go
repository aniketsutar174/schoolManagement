package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type StudDatabase struct {
	Id    primitive.ObjectID `json:"_id,omitempty" bson:"_id, omitempty"`
	Name  string             `json:"name,omitempty" bson:"name, omitempty"`
	Roll  int                `json:"roll,omitempty" bson:"roll, omitempty"`
	Class int                `json:"class,omitempty" bson:"class, omitempty"`
	Marks Marks              `json:"marks,omitempty" bson:"marks, omitempty"`
}

type Marks struct {
	Maths   float32 `json:"maths,omitempty" bson:"maths, omitempty"`
	Science float32 `json:"science,omitempty" bson:"science, omitempty"`
}
