package model

import "TeacherKPI/database"

//Этап выполнения задачи (Task)

type Stage struct {
	TaskId      uint64 `json:"taskId"`
	Owner       Task   `json:"-"`
	StageTypeId uint64 `json:"typeId"` //Индекс типа этапа
	Status      int    `json:"status"` //Текущий статус
	QueuePos    int    `json:"-"`      //Позиция в списке этапов (привязана к типу этапа)
}

const (
	STAGE_OPEN      int = 0
	STAGE_CLOSED    int = 1
	STAGE_UNCHECKED int = 2
)

func StageFromParamsMap(values map[string]interface{}, useLinkedEntities bool) Stage {
	ret := Stage{
		TaskId:      values["task_id"].(uint64),
		StageTypeId: values["stage_id"].(uint64),
		Status:      values["status"].(int),
		QueuePos:    values["position"].(int),
	}

	if useLinkedEntities {
		userValues := database.GetTaskById(ret.TaskId)
		ret.Owner = TaskFromParamsMap(userValues, true)
	}

	return ret
}

func (stage Stage) SetStatus(status int) {
	stage.Status = status

	//Если выполнена задача, то предыдущие также выполнены
	if status == STAGE_CLOSED {
		for _, otherStage := range stage.Owner.Stages {
			if otherStage.QueuePos < stage.QueuePos {
				otherStage.Status = STAGE_CLOSED
			} else {
				break
			}
		}
	}
}
