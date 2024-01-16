package model

import "database/sql"

//BRANCHED

//Пока эта структура просто для ее помещения в связанные сущности
//Реализуем потом

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	TasksIds []int  `json:"-"`
	//TODO
}

func (user *User) GetName() string {
	return user.Username
}
func (user *User) SetName(username string) {
	user.Username = username
}

func (user *User) ScanFromRows(rows *sql.Rows) error {
	ret := new(User)
	if err := rows.Scan(&ret.Id, &ret.Username); err != nil {
		return err
	}
	return nil
}
