package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
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

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	db := getConnection()

	fmt.Println(r)

	requestDto := CreatePostRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	q, err := db.Prepare("INSERT INTO post (title, content) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = q.Exec(requestDto.Title, requestDto.Content)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func main() {
	http.HandleFunc("/post/create", CreatePost)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
