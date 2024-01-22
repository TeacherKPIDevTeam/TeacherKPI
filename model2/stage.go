package model2

//Этап выполнения задачи (Task)

type Stage struct {
	Id         int    `json:"id"`
	TaskId     int    `json:"taskId"`
	TypeId     int    `json:"typeId"` //Индекс типа этапа
	TypeName   string `json:"typeName"`
	Status     int    `json:"status"` //Текущий статус
	StatusName string `json:"statusName"`
	QueuePos   int    `json:"-"` //Позиция в списке этапов (привязана к типу этапа)
}

var stagesCache = map[uint64]*Stage{}

const (
	STAGE_OPEN      int = 0
	STAGE_CLOSED    int = 1
	STAGE_UNCHECKED int = 2
)

//
//CONVERTION DATA FROM DB INTO ENTITY
//

// Преобразовывает данные из базы, полученные через database.GetStageById() в
// структуру. Вызывается из StageById в случае, если Stage с нужным id еще не кеширован

//func StageFromParamsMap(values map[string]interface{}) *Stage {
//	ret := Stage{
//		Id:         values["id"].(uint64),
//		TaskId:     values["task_id"].(uint64),
//		TypeId:     values["type_id"].(uint64),
//		TypeName:   values["type_name"].(string),
//		Status:     values["status"].(int),
//		QueuePos:   values["queue_pos"].(int),
//		StatusName: values["status_name"].(string),
//	}

//	return &ret
//}

/*func (stage *Stage) SetStatus(status int) {
	stage.Status = status

	//Если выполнена задача, то предыдущие также выполнены
	if status == STAGE_CLOSED {
		parentTaskStages, _ := StagesByTaskId(stage.TaskId)
		for _, otherStage := range parentTaskStages {
			if otherStage.QueuePos < stage.QueuePos {
				otherStage.Status = STAGE_CLOSED
				otherStage.Save()
			} else {
				break
			}
		}
	}
}*/

//
//API METHODS
//
/*
func HttpGetStageById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)
	stage, err := StageById(id)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(stage)
}

// HTTP GET STAGE FOR TASK
func HttpGetStagesByTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var owner string
	if r.URL.Path == "/stages" {
		owner = r.URL.Query().Get("owner")
	}
	if strings.HasPrefix(r.URL.Path, "/task/") {
		owner = mux.Vars(r)["id"]
	}

	if owner != "" {
		ownerId, _ := strconv.ParseUint(owner, 10, 64)
		stages, err := StagesByTaskId(ownerId)

		if err != nil {
			w.WriteHeader(404)
			return
		}

		json.NewEncoder(w).Encode(stages)
	}
}*/
