package service

import (
	"encoding/json"
	"fmt"
	"github.com/KimTaeO/GO-CRUD/config"
	presentation "github.com/KimTaeO/GO-CRUD/post/presentation/dto/request"
	"net/http"
	"strconv"
)

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}

	db := config.GetConnection()

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(id)

	requestDto := presentation.UpdatePostRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	q, err := db.Prepare("UPDATE post SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = q.Exec(requestDto.Title, requestDto.Content, id)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
