package repository

import (
	"github.com/KimTaeO/GO-CRUD/config"
	"github.com/KimTaeO/GO-CRUD/post/entity"
)

func Save(post entity.Post) {
	db := config.GetConnection()

	q, err := db.Prepare("INSERT INTO post (title, content) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = q.Exec(post.Title, post.Content)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
