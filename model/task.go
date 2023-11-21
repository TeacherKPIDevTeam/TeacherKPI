package model

type Task struct {
	Id uint64 `json:"id"`
	//Какому пользователю принадлежит задача (User)
	UserId uint64 `json:"user_id"`
	Owner  *User
	//Индекс типа задач
	TaskTypeId uint64 `json:"type_id"`
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
