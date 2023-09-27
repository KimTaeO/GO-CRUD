package service

import (
	"encoding/json"
	"github.com/KimTaeO/GO-CRUD/config"
	presentation "github.com/KimTaeO/GO-CRUD/post/presentation/dto/request"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}

	db := config.GetConnection()

	requestDto := presentation.CreatePostRequest{}
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
