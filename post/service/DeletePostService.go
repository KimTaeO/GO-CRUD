package service

import (
	"github.com/KimTaeO/GO-CRUD/config"
	"net/http"
	"strconv"
)

func DeletePostById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}

	db := config.GetConnection()

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}

	q, err := db.Prepare("DELETE FROM post WHERE id = ?")

	if _, err := q.Exec(id); err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
