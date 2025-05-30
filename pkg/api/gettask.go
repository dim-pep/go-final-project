package api

import (
	"fmt"
	"go1f/pkg/db"
	"net/http"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, map[string]string{"error": fmt.Sprintf("ошибка получения задачи: %v", err)}, http.StatusInternalServerError)
		return
	}
	writeJson(w, task, http.StatusOK)
}
