package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDb() error {
	if db != nil {
		return nil
	}
	var err error
	db, err = sql.Open("mysql", "root@tcp(localhost:3306)/teacher_kpi")
	if err != nil {
		//Добавить обработчик отсутствия подключения к базе
		//panic(err)
		return err
	}
	return nil
}

func RequestQuery(request string, args ...any) (*sql.Rows, error) {
	/*if err := InitDb(); err != nil {
		return nil, err
	}*/

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/teacher_kpi")
	if err != nil {
		return nil, err
	}

	result, err := db.Query(request, args...)

	if err != nil {
		//panic(err)
		db.Close()
		return nil, err
	}

	db.Close()
	return result, nil
}

func RequestNonQuery(request string, args ...any) (sql.Result, error) {
	/*if err := InitDb(); err != nil {
		return err
	}*/

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/teacher_kpi")
	if err != nil {
		return nil, err
	}

	result, err := db.Exec(request, args...)

	if err != nil {
		//panic(err)
		db.Close()
		return nil, err
	}

	db.Close()
	return result, nil
}
