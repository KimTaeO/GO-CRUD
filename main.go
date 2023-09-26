package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func getConnection() *sql.DB {
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

func main() {
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
