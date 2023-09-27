package service

import (
	"encoding/json"
	"github.com/KimTaeO/GO-CRUD/post/entity"
	presentation "github.com/KimTaeO/GO-CRUD/post/presentation/dto/request"
	"github.com/KimTaeO/GO-CRUD/post/repository"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}

	requestDto := presentation.CreatePostRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	post := entity.Post{
		Title:   requestDto.Title,
		Content: requestDto.Content,
	}

	repository.Save(post)

}
