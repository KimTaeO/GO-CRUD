package service

import (
	"github.com/KimTaeO/GO-CRUD/post/repository"
	"net/http"
	"strconv"
)

func DeletePostById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}

	repository.DeleteById(id)
}
