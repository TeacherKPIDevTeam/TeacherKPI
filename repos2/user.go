package repos2

import (
	"TeacherKPI/database"
	"TeacherKPI/model"
)

type UserRepo struct {
	Db *database.DbConnection
}

func (repo UserRepo) Create(user *model.User) error {
	result, err := repo.Db.RequestNonQuery("INSERT INTO users VALUES(NULL, ?)", user.Username)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	user.Id = int(id)
	return nil
}

func (repo UserRepo) GetById(id int) (*model.User, error) {
	response, err := repo.Db.RequestQuery("SELECT * FROM users WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	response.Next()
	defer response.Close()

	user := new(model.User)
	if err := user.ScanFromRows(response); err != nil {
		return nil, err
	}
	user.TasksIds, _ = repo.GetUserTaskIds(user.Id)
	return user, nil
}

func (repo UserRepo) GetUserTasks(user *model.User) ([]*model.Task, error) {
	response, err := repo.Db.RequestQuery("SELECT * FROM tasks WHERE user_id=?", user.Id)
	if err != nil {
		return nil, err
	}

	response.Next()
	defer response.Close()

	ret := []*model.Task{}

	for {
		if !response.Next() {
			break
		}

		task := new(model.Task)
		if err := response.Scan(&task.Id, &task.UserId, &task.TypeId); err != nil {
			return nil, err
		}

		ret = append(ret, task)
	}

	return ret, nil
}

func (repo UserRepo) GetUserTaskIds(userId int) ([]int, error) {
	response, err := repo.Db.RequestQuery("SELECT * FROM tasks WHERE user_id=?", userId)
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

func (repo UserRepo) SearchByUsername(username string) ([]*model.User, error) {
	response, err := repo.Db.RequestQuery("SELECT * FROM users WHERE name LIKE '%?%'", username)
	if err != nil {
		return nil, err
	}
	defer response.Close()

	ret := []*model.User{}
	for {
		if !response.Next() {
			break
		}

		user := model.User{}
		if err := response.Scan(&user.Id, &user.Username); err != nil {
			return nil, err
		}

		ret = append(ret, &user)
	}

	return ret, nil
}

func (repo UserRepo) AddTask(user *model.User, task *model.Task) {
	for _, v := range user.TasksIds {
		if v == task.Id {
			return
		}
	}
	user.TasksIds = append(user.TasksIds, task.Id)
	repo.Update(user)
}

func (repo UserRepo) Update(user *model.User) error {
	_, err := repo.Db.RequestNonQuery("UPDATE users SET username = ? WHERE id = ?",
		user.Username,
		user.Id)
	if err != nil {
		return err
	}
	return err
}

func (repo UserRepo) Delete(user *model.User) error {
	//TODO
	return nil
}
