package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetPostResponse struct {
	Id      string `json:"id"`
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
		http.Error(w, "Forbidden", http.StatusForbidden)
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

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}

	db := getConnection()

	fmt.Println(r)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(id)

	requestDto := UpdatePostRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	q, err := db.Prepare("UPDATE post SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = q.Exec(requestDto.Title, requestDto.Content, id)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func GetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}

	db := getConnection()

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}

	rows := db.QueryRow("SELECT id, title, content FROM post WHERE id = ?", id)

	response := GetPostResponse{}
	if err := rows.Scan(&response.Id, &response.Title, &response.Content); err != nil {
		panic(err.Error())
	}

	fmt.Println(response)

	serialized, err := json.Marshal(response)
	if err != nil {
		panic(err.Error())
	}

	_, err = w.Write(serialized)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func main() {
	http.HandleFunc("/create", CreatePost)
	http.HandleFunc("/read", GetById)
	http.HandleFunc("/update", UpdatePost)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
