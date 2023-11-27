package model

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct {
	Id uint64 `json:"Id"`
	//Какому пользователю принадлежит задача (User)
	UserId uint64 `json:"UserId"`
	Owner  *User
	//Индекс типа задач
	TaskTypeId uint64 `json:"Type"`
	//Этапы выполнения
	Stages []*Stage
}

func (task *Task) AddStage(stage Stage) {
	stage.Owner = task
	stagePos := stage.QueuePos
	for k, v := range task.Stages {
		//Вставляем stage соответственно порядковому номеру
		if v.QueuePos > stagePos {
			task.Stages = append(task.Stages[:k+1], task.Stages[k:]...)
			task.Stages[k] = &stage
			return
		}
	}
	//Если добавляемая задача по позиции последняя - просто добавляем в конец
	task.Stages = append(task.Stages, &stage)
}

//
//API METHODS
//

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	owner := r.URL.Query().Get("owner")

	if owner != "" {
		user := User{Username: owner}
		json.NewEncoder(w).Encode(Task{Owner: &user})
	}
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)
	json.NewEncoder(w).Encode(Task{Id: id})
}
