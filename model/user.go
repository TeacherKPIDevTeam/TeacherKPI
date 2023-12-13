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
	UserId   uint64 `json:"id"`
	Username string `json:"username"`
	Tasks    []Task `json:"-"`
	//TODO
}

func UserFromParamsMap(values map[string]interface{}, useLinkedEntities bool) User {
	ret := User{
		UserId:   values["id"].(uint64),
		Username: values["username"].(string),
	}

	if useLinkedEntities {
		ret.Tasks = TasksByUserId(ret.UserId)
	}

	return ret
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	values := database.GetUserById(id)
	if values == nil {
		json.NewEncoder(w).Encode("Error: not found")
		return
	}

	user := UserFromParamsMap(values, false)
	json.NewEncoder(w).Encode(user)
}
