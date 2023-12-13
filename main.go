package main

import (
	"errors"
	"log"
	"net/http"

	"TeacherKPI/model"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", model.GetTasksByUser).Methods("GET")
	router.HandleFunc("/user/{id}/tasks", model.GetTasksByUser).Methods("GET")
	router.HandleFunc("/task/{id}", model.GetTaskById).Methods("GET")
	router.HandleFunc("/user/{id}", model.GetUserById).Methods("GET")
	//r.HandleFunc("/task", createBook).Methods("POST")
	//r.HandleFunc("/task/{id}", updateBook).Methods("PUT")
	//r.HandleFunc("/task/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func div(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero!")
	}
	return a / b, nil
}
