package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func main() {
	http.HandleFunc("/student", getStudentData)
	http.ListenAndServe(":8080", nil)
}

func getStudentData(w http.ResponseWriter, r *http.Request) {
	// Call the other microservice API to get student data
	resp, err := http.Get("http://localhost:4000/stud/allData")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse the JSON response and extract the student data
	var students []StudDatabase
	err = json.Unmarshal(body, &students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the student data as JSON
	jsonBytes, err := json.Marshal(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
