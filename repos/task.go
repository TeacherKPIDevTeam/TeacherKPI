package repos

import (
	"TeacherKPI/database"
	"TeacherKPI/model"
)

type TaskRepo struct {
	Db *database.DbConnection
}

func (repo TaskRepo) Create(task *model.Task) error {
	result, err := repo.Db.RequestNonQuery("INSERT INTO tasks VALUES(NULL, ?, ?)", task.UserId, task.TypeId)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	task.Id = int(id)
	return nil
}

func (repo TaskRepo) GetById(id int) (*model.Task, error) {
	response, err := repo.Db.RequestQuery("SELECT * FROM tasks WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	response.Next()
	defer response.Close()

	task := new(model.Task)
	if err := response.Scan(&task.Id, &task.UserId, &task.TypeId); err != nil {
		return nil, err
	}
	return task, nil
}

func (repo TaskRepo) GetTaskStages(task *model.Task) ([]*model.Stage, error) {
	response, err := repo.Db.RequestQuery("SELECT * FROM tasks WHERE user_id=?", task.Id)
	if err != nil {
		return nil, err
	}

	response.Next()
	defer response.Close()

	ret := []*model.Stage{}

	for {
		if !response.Next() {
			break
		}

		stage := new(model.Stage)
		if err := response.Scan(&stage.Id); err != nil {
			return nil, err
		}

		ret = append(ret, stage)
	}

	return ret, nil
}

func (repo TaskRepo) GetTaskStagesIds(taskId int) ([]int, error) {
	response, err := repo.Db.RequestQuery("SELECT * FROM stage WHERE user_id=?", taskId)
	if err != nil {
		return nil, err
	}

	response.Next()
	defer response.Close()

	ret := []int{}
	var tempId int
	for {
		if !response.Next() {
			break
		}

		if err := response.Scan(&tempId); err != nil {
			return nil, err
		}

		ret = append(ret, tempId)
	}
	return ret, nil
}

func (repo TaskRepo) AddStage(task *model.Task, stage *model.Stage) {
	for _, v := range task.StagesIds {
		if v == task.Id {
			return
		}
	}
	task.StagesIds = append(task.StagesIds, stage.Id)
	repo.Update(task)
}

func (repo TaskRepo) Update(task *model.Task) error {
	_, err := repo.Db.RequestNonQuery("UPDATE tasks SET user_id=?, type_id=? WHERE id=?",
		task.UserId,
		task.TypeId,
		task.Id)
	if err != nil {
		return err
	}
	return err
}

func (repo TaskRepo) Delete(task *model.Task) error {
	//TODO
	return nil
}
