package api

import (
	"go1f/pkg/db"
	"net/http"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(30)
	if err != nil {
		writeJson(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}
	if tasks == nil {
		writeJson(w, TasksResp{Tasks: []*db.Task{}}, http.StatusOK)
	} else {
		writeJson(w, TasksResp{Tasks: tasks}, http.StatusOK)
	}
}

//s
