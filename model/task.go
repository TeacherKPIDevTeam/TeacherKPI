package model

import "database/sql"

type Task struct {
	Id     int `json:"id"`
	UserId int `json:"userId"`
	//Индекс типа задач
	TypeId    int    `json:"type"`
	TypeName  string `json:"typeName"`
	StagesIds []int  `json:"stagesIds"`
}

func (task *Task) ScanFromRows(rows *sql.Rows) error {
	ret := new(Task)
	if err := rows.Scan(&ret.Id, &ret.UserId, &ret.TypeId); err != nil {
		return err
	}
	return nil
}
