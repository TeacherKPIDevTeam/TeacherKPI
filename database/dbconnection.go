package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DbConnection struct {
	Settings DbConnectionSettings
}

func (con DbConnection) InitConnection() (*sql.DB, error) {
	var err error

	db, err := sql.Open(con.Settings.Dbms, con.Settings.ConString())
	if err != nil {
		//Добавить обработчик отсутствия подключения к базе
		//panic(err)
		return nil, err
	}
	return db, nil
}

func (con DbConnection) RequestQuery(request string, args ...any) (*sql.Rows, error) {
	db, err := con.InitConnection()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	result, err := db.Query(request, args...)

	if err != nil {
		//panic(err)
		return nil, err
	}

	return result, nil
}

func (con DbConnection) RequestNonQuery(request string, args ...any) (sql.Result, error) {
	db, err := con.InitConnection()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	result, err := db.Exec(request, args...)

	if err != nil {
		//panic(err)
		return nil, err
	}

	return result, nil
}
