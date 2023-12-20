package tests

import (
	"testing"
)

func Test_stage(t *testing.T) {
	/*task := model.Task{}

	//Тест присвоения ownerа при AddStage
	task.AddStage(&model.Stage{QueuePos: 1, Status: model.STAGE_OPEN})
	assert.ObjectsAreEqual(task.Stages[0].Task, task)

	//Тест добавления стадий и адекватности их сортировки
	task.AddStage(&model.Stage{QueuePos: 2, Status: model.STAGE_OPEN})
	task.AddStage(&model.Stage{QueuePos: 0, Status: model.STAGE_OPEN})
	task.AddStage(&model.Stage{QueuePos: 4, Status: model.STAGE_OPEN})
	task.AddStage(&model.Stage{QueuePos: 3, Status: model.STAGE_OPEN})
	order := [5]int{}
	expectedOrder := [5]int{0, 1, 2, 3, 4}
	for k, v := range task.Stages {
		order[k] = v.QueuePos
	}
	assert.Equal(t, expectedOrder, order)

	//Проверка присвоения статуса CLOSED всем задачам до текущей
	task.Stages[2].SetStatus(model.STAGE_CLOSED)
	statuses := [3]int{task.Stages[0].Status, task.Stages[1].Status, task.Stages[2].Status}
	expectedStatuses := [3]int{
		model.STAGE_CLOSED,
		model.STAGE_CLOSED,
		model.STAGE_CLOSED,
	}
	assert.Equal(t, expectedStatuses, statuses)*/
}
