package database2

import "fmt"

type DbConnectionSettings struct {
	Username string
	Password string
	Host     string
	Port     uint16
	DbName   string
	Dbms     string
}

func DbConnectionSettings_New() DbConnectionSettings {
	return DbConnectionSettings{
		Username: "root",
		Password: "",
		Host:     "localhost",
		Port:     3306,
		DbName:   "",
		Dbms:     "mysql",
	}
}

func (s DbConnectionSettings) ConString() string {
	password := s.Password
	if password != "" {
		password = ":" + password
	}
	return fmt.Sprintf("%s%s@tcp(%s:%d)/%s", s.Username, password, s.Host, s.Port, s.DbName)
}
