package service

import (
	"encoding/json"
	dto "github.com/KimTaeO/GO-CRUD/post/presentation/dto/response"
	"github.com/KimTaeO/GO-CRUD/post/repository"
	"net/http"
	"strconv"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}

	post, err := repository.GetById(id)
	if err != nil {
		http.Error(w, "post not found", http.StatusNotFound)
	}

	response := dto.GetPostResponse{
		Id:      post.Id,
		Title:   post.Title,
		Content: post.Content,
	}

	serialized, err := json.Marshal(response)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(serialized)
	if err != nil {
		panic(err.Error())
	}
}
