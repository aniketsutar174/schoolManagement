package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aniketsutar174/schoolManagement/model"
	"github.com/aniketsutar174/schoolManagement/service"
	"github.com/gorilla/mux"
)

func AllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var studData []model.StudDatabase
	studentDataReturn, err := service.AllDataStud(studData)

	if err != nil {
		log.Fatal("unable to show")
	}
	json.NewEncoder(w).Encode(studentDataReturn)
}

func AllDataById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := mux.Vars(r)["id"]

	studDataById, _ := service.AllDataByIdStud(id)

	json.NewEncoder(w).Encode(studDataById)
}

func UpdateStudDataById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	id := params["id"]

	log.Println(id)
	var studData model.StudDatabase
	err := json.NewDecoder(r.Body).Decode(&studData)
	if err != nil {
		log.Fatal("unable to show")
	}

	result, err := service.UpdateStudDataByIdStud(id, studData)
	if err != nil {
		log.Fatal("unable to show")
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
func InsertData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var studData model.StudDatabase
	err := json.NewDecoder(r.Body).Decode(&studData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := service.InsertDataStud(studData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func DeleteDataById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	id := params["id"]

	result, err := service.DeleteDataById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(result)
}
