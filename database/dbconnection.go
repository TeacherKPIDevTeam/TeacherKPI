package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//var db *sql.DB

func RequestQuery(request string, args ...any) *sql.Rows {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/teacher_kpi")

	if err != nil {
		//Добавить обработчик отсутствия подключения к базе
		panic(err)
	}
	defer db.Close()

	result, err := db.Query(request, args...)

	if err != nil {
		//panic(err)
	}

	return result
}

func RequestNonQuery(request string, args ...any) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/teacher_kpi")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(request, args)

	if err != nil {
		panic(err)
	}
}
