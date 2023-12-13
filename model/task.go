package model

import (
	"TeacherKPI/database"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Task struct {
	Id uint64 `json:"Id"`
	//Какому пользователю принадлежит задача (User)
	UserId uint64 `json:"UserId"`
	Owner  User   `json:"-"`
	//Индекс типа задач
	TaskTypeId uint64 `json:"Type"`
	//Этапы выполнения
	Stages []Stage `json:"-"`
}

func TaskFromParamsMap(values map[string]interface{}, useLinkedEntities bool) Task {
	ret := Task{
		Id:         values["id"].(uint64),
		UserId:     values["owner_id"].(uint64),
		TaskTypeId: values["type_id"].(uint64),
	}

	if useLinkedEntities {
		userValues := database.GetUserById(ret.UserId)
		ret.Owner = UserFromParamsMap(userValues, true)
	}

	return ret
}

func TasksByUserId(userId uint64) []Task {
	ret := []Task{}
	tasksValues := database.GetTasksByUserId(userId)
	for _, v := range tasksValues {
		task := TaskFromParamsMap(v.(map[string]interface{}), false)
		ret = append(ret, task)
	}
	return ret
}

func (task Task) AddStage(stage Stage) {
	stage.Owner = task
	stagePos := stage.QueuePos
	for k, v := range task.Stages {
		//Вставляем stage соответственно порядковому номеру
		if v.QueuePos > stagePos {
			task.Stages = append(task.Stages[:k+1], task.Stages[k:]...)
			task.Stages[k] = stage
			return
		}
	}
	//Если добавляемая задача по позиции последняя - просто добавляем в конец
	task.Stages = append(task.Stages, stage)
}

//
//API METHODS
//

func GetTasksByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var owner string
	if r.URL.Path == "/tasks" {
		owner = r.URL.Query().Get("owner")
	}
	if strings.HasPrefix(r.URL.Path, "/user/") {
		owner = mux.Vars(r)["id"]
	}

	if owner != "" {
		ownerId, _ := strconv.ParseUint(owner, 10, 64)
		tasks := TasksByUserId(ownerId)

		json.NewEncoder(w).Encode(tasks)
	}
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)
	values := database.GetTaskById(id)

	if values == nil {
		w.WriteHeader(404)
		return
	}

	task := TaskFromParamsMap(values, false)
	json.NewEncoder(w).Encode(task)
}
