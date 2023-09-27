package service

import (
	"encoding/json"
	"github.com/KimTaeO/GO-CRUD/post/entity"
	presentation "github.com/KimTaeO/GO-CRUD/post/presentation/dto/request"
	"github.com/KimTaeO/GO-CRUD/post/repository"
	"net/http"
	"strconv"
)

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}

	requestDto := presentation.UpdatePostRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	post := entity.Post{
		Title:   requestDto.Title,
		Content: requestDto.Content,
	}

	repository.UpdateById(id, post)
}
