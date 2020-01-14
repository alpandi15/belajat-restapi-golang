package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	host     = "127.0.0.1:3306"
	username = "root"
	password = "111111"
	dbname   = "db_golang"
	charset  = "utf8"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+")/"+dbname)
	if err != nil {
		return nil, err
	}

	return db, nil
}
