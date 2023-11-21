package main

import (
	"errors"
	"fmt"

	"github.com/TeacherKPIDevTeam/TeacherKPI/model"
)

func main() {
	testTask := model.Task{Id: 0}
	testTask.AddStage(model.Stage{QueuePos: 1})
	testTask.AddStage(model.Stage{QueuePos: 2})
	testTask.AddStage(model.Stage{QueuePos: 0})
	testTask.AddStage(model.Stage{QueuePos: 4})
	testTask.AddStage(model.Stage{QueuePos: 3})

	divResult, _ := div(5, 2)
	fmt.Println(divResult, "Hello, world!")
}

func div(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero!")
	}
	return a / b, nil
}
