package config

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	config := mysql.Config{
		User:   "root",
		Passwd: "1234",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "go",
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("connected")
	return db
}
