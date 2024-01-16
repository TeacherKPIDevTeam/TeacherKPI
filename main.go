package main

import (
	"TeacherKPI/api"
	"TeacherKPI/database"
	"TeacherKPI/repos"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	dbConnection := new(database.DbConnection)
	dbConnection.Settings = database.DbConnectionSettings_New()
	dbConnection.Settings.DbName = "teacher_kpi"

	userApi := &api.UserApi{
		Users: &repos.UserRepo{
			Db: dbConnection,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", userApi.HttpGetUserById).Methods("GET")
	//router.HandleFunc("/user/{id}/tasks", model.HttpGetTasksByUser).Methods("GET")
	//router.HandleFunc("/tasks", model.HttpGetTasksByUser).Methods("GET")
	//router.HandleFunc("/task/{id}", model.HttpGetTaskById).Methods("GET")
	//router.HandleFunc("/task/{id}/stages", model.HttpGetStagesByTask).Methods("GET")
	//router.HandleFunc("/stages", model.HttpGetStagesByTask).Methods("GET")
	//router.HandleFunc("/stage/{id}", model.HttpGetStageById).Methods("GET")

	//r.HandleFunc("/task", createBook).Methods("POST")
	//r.HandleFunc("/task/{id}", updateBook).Methods("PUT")
	//r.HandleFunc("/task/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8001", router))
}
