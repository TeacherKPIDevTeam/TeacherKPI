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
	Id     uint64 `json:"id"`
	UserId uint64 `json:"userId"`
	//Индекс типа задач
	TypeId    uint64   `json:"type"`
	TypeName  string   `json:"typeName"`
	StagesIds []uint64 `json:"stagesIds"`
}

var tasksCache = map[uint64]*Task{}

// Преобразовывает данные из базы, полученные через database.GetTaskById() в
// структуру. Вызывается из TaskById в случае, если Task с нужным id еще не кеширован
func TaskFromParamsMap(values map[string]interface{}) *Task {
	ret := Task{
		Id:     values["id"].(uint64),
		UserId: values["owner_id"].(uint64),
		TypeId: values["type_id"].(uint64),
	}
	ret.StagesIds, _ = database.GetStageIdsByTaskId(ret.Id)

	return &ret
}

func TaskById(id uint64) (*Task, error) {
	if _, exists := tasksCache[id]; !exists {
		values, err := database.GetTaskDataById(id)
		if err != nil {
			return nil, err
		}
		tasksCache[id] = TaskFromParamsMap(values)
	}
	return tasksCache[id], nil
}

func TasksByUserId(userId uint64) ([]*Task, error) {
	ret := []*Task{}
	tasksValues, err := database.GetTaskIdsByUserId(userId)
	if err != nil {
		return nil, err
	}
	for _, id := range tasksValues {
		task, err := TaskById(id)
		if err != nil {
			return nil, err
		}
		ret = append(ret, task)
	}
	return ret, nil
}

// SAVE ENTITY INTO DB
func (task *Task) Save() {

}

//
//FUNCTIONAL METHODS
//

func (task *Task) AddStage(stage *Stage) {
	stage.TaskId = task.Id
	stage.Save()

	/*stagePos := stage.QueuePos

	/*taskStages, _ := StagesByTaskId(task.Id)
	for k, v := range taskStages {
		//Вставляем stage соответственно порядковому номеру
		if v.QueuePos > stagePos {
			task.Stages = append(task.Stages[:k+1], task.Stages[k:]...)
			task.Stages[k] = stage
			return
		}
	}
	//Если добавляемая задача по позиции последняя - просто добавляем в конец
	task.Stages = append(task.Stages, stage)*/
}

// HTTP GET TASK BY ID
func HttpGetTaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)
	task, err := TaskById(id)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// HTTP GET TASKS FOR USER
func HttpGetTasksByUser(w http.ResponseWriter, r *http.Request) {
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
		tasks, err := TasksByUserId(ownerId)

		if err != nil {
			w.WriteHeader(404)
			return
		}

		json.NewEncoder(w).Encode(tasks)
	}
}
