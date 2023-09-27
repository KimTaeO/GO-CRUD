package controller

import (
	"github.com/KimTaeO/GO-CRUD/post/service"
	"net/http"
)

func PostController() {
	http.HandleFunc("/create", service.CreatePost)
	http.HandleFunc("/read", service.CreatePost)
	http.HandleFunc("/update", service.UpdatePost)
	http.HandleFunc("/delete", service.DeletePostById)
}
