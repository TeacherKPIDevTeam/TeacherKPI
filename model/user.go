package model

import (
	"encoding/json"
	"net/http"
	"strconv"

	"TeacherKPI/database"

	"github.com/gorilla/mux"
)

//Пока эта структура просто для ее помещения в связанные сущности
//Реализуем потом

type User struct {
	Id       uint64   `json:"id"`
	Username string   `json:"username"`
	TasksIds []uint64 `json:"-"`
	//TODO
}

var usersCache = map[uint64]*User{}

/*

	РАБОТА С БД

*/

// Преобразовывает данные из базы, полученные через database.GetUserById() в
// структуру. Вызывается из UserById в случае, если User с нужным id еще не кеширован
func UserFromParamsMap(values map[string]interface{}) *User {
	ret := User{
		Id:       values["id"].(uint64),
		Username: values["username"].(string),
	}
	ret.TasksIds, _ = database.GetTaskIdsByUserId(ret.Id)

	return &ret
}

func (user *User) ToParamsMap() map[string]interface{} {
	ret := map[string]interface{}{
		"id":       user.Id,
		"username": user.Username,
	}

	return ret
}

// Метод получения User по id. Пытается извлечь из кеша, если не выходит - обращается к БД
func UserById(id uint64) (*User, error) {
	if _, exists := usersCache[id]; !exists {
		values, err := database.GetUserDataById(id)
		if err != nil {
			return nil, err
		}
		usersCache[id] = UserFromParamsMap(values)
	}
	return usersCache[id], nil
}

// Создает объект пользователя и сохраняет в базу
func CreateUser(username string) *User {
	ret := User{
		Username: username,
	}
	result, _ := database.CreateUser(ret.ToParamsMap())
	id, _ := result.LastInsertId()
	ret.Id = uint64(id)
	usersCache[ret.Id] = &ret
	return &ret
}

// Обновляет пользователя в базе
func (user *User) Save() {
	database.SaveUser(user.ToParamsMap())
}

/*

	ГЕТТЕРЫ СЕТТЕРЫ

*/

func (user *User) GetName() string {
	return user.Username
}
func (user *User) SetName(username string) {
	user.Username = username
	user.Save()
}

func (user *User) AddTask(task *Task) {
	task.UserId = user.Id
	user.TasksIds = append(user.TasksIds, task.Id)
	task.Save()
}

func (user *User) Tasks() []*Task {
	tasks, _ := TasksByUserId(user.Id)
	return tasks
}

// HTTP GET
func HttpGetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	user, err := UserById(id)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func HttpPostUser(w http.ResponseWriter, r *http.Request) {
	/*w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	user, err := UserById(id)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(user)*/
}
