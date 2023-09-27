package service

import (
	"encoding/json"
	"github.com/KimTaeO/GO-CRUD/config"
	presentation2 "github.com/KimTaeO/GO-CRUD/post/presentation/dto/response"
	"net/http"
	"strconv"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}

	db := config.GetConnection()

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}

	rows := db.QueryRow("SELECT id, title, content FROM post WHERE id = ?", id)

	response := presentation2.GetPostResponse{}
	if err := rows.Scan(&response.Id, &response.Title, &response.Content); err != nil {
		http.Error(w, "post not found", http.StatusNotFound)
		return
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

	defer db.Close()
}
