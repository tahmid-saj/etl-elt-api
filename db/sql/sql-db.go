package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var SqlDB *sql.DB

func InitSQLDB() {
	SqlDB, err := sql.Open("sqlite3", "etl_elt_api.db")

	if err != nil {
		panic("could not connect to the sql database")
	}

	SqlDB.SetConnMaxIdleTime(5)
	SqlDB.SetMaxOpenConns(10)
}