package main

import (
	"log"
	"net/http"

	"TeacherKPI/model"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", model.HttpGetUserById).Methods("GET")
	router.HandleFunc("/user/{id}/tasks", model.HttpGetTasksByUser).Methods("GET")
	router.HandleFunc("/tasks", model.HttpGetTasksByUser).Methods("GET")
	router.HandleFunc("/task/{id}", model.HttpGetTaskById).Methods("GET")
	router.HandleFunc("/task/{id}/stages", model.HttpGetStagesByTask).Methods("GET")
	router.HandleFunc("/stages", model.HttpGetStagesByTask).Methods("GET")
	router.HandleFunc("/stage/{id}", model.HttpGetStageById).Methods("GET")

	//r.HandleFunc("/task", createBook).Methods("POST")
	//r.HandleFunc("/task/{id}", updateBook).Methods("PUT")
	//r.HandleFunc("/task/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
