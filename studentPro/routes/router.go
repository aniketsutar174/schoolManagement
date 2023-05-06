package routes

import (
	"github.com/aniketsutar174/schoolManagement/handler"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/stud/allData", handler.AllData).Methods("GET")
	router.HandleFunc("/stud/allData/{id}", handler.AllDataById).Methods("GET")
	router.HandleFunc("/stud/updateStudData/{id}", handler.UpdateStudDataById).Methods("PUT")
	router.HandleFunc("/stud/deleteData/{id}", handler.DeleteDataById).Methods("DELETE")
	router.HandleFunc("/stud/insertData", handler.InsertData).Methods("POST")

	return router
}
