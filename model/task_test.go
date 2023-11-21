package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stage(t *testing.T) {
	task := Task{}

	//Тест присвоения ownerа при AddStage
	task.AddStage(Stage{QueuePos: 1, Status: STAGE_OPEN})
	assert.ObjectsAreEqual(task.Stages[0].Owner, task)

	//Тест добавления стадий и адекватности их сортировки
	task.AddStage(Stage{QueuePos: 2, Status: STAGE_OPEN})
	task.AddStage(Stage{QueuePos: 0, Status: STAGE_OPEN})
	task.AddStage(Stage{QueuePos: 4, Status: STAGE_OPEN})
	task.AddStage(Stage{QueuePos: 3, Status: STAGE_OPEN})
	order := [5]int{}
	expectedOrder := [5]int{0, 1, 2, 3, 4}
	for k, v := range task.Stages {
		order[k] = v.QueuePos
	}
	assert.Equal(t, expectedOrder, order)

	//Проверка присвоения статуса CLOSED всем задачам до текущей
	task.Stages[2].SetStatus(STAGE_CLOSED)
	statuses := [3]int{task.Stages[0].Status, task.Stages[1].Status, task.Stages[2].Status}
	expectedStatuses := [3]int{STAGE_CLOSED, STAGE_CLOSED, STAGE_CLOSED}
	assert.Equal(t, expectedStatuses, statuses)
}
