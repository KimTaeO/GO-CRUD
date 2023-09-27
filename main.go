package main

import (
	controller "github.com/KimTaeO/GO-CRUD/post/presentation"
	"log"
	"net/http"
)

func main() {
	controller.PostController()
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
