package model

//Этап выполнения задачи (Task)

type Stage struct {
	TaskId      uint64
	Owner       *Task
	StageTypeId uint64 //Индекс типа этапа
	Status      int    //Текущий статус
	QueuePos    int    //Позиция в списке этапов (привязана к типу этапа)
}

const (
	STAGE_OPEN      int = 0
	STAGE_CLOSED    int = 1
	STAGE_UNCHECKED int = 2
)

func (stage *Stage) SetStatus(status int) {
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
