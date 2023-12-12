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
	UserId   uint64
	Username string
	//TODO
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	uid, username := database.GetUserById(id)
	user := User{
		UserId:   uid,
		Username: username,
	}

	json.NewEncoder(w).Encode(user)
}
