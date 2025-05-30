package api

import (
	"fmt"
	"go1f/pkg/db"
	"net/http"
	"time"
)

func doneHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, map[string]string{"error": fmt.Sprintf("ошибка получения задачи: %v", err)}, http.StatusInternalServerError)
		return
	}
	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeJson(w, map[string]string{"error": fmt.Sprintf("ошибка удаления задачи при выполнении: %v", err)}, http.StatusInternalServerError)
			return
		}
		writeJson(w, map[string]int64{}, http.StatusOK)
	} else {
		nextDate, err := NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			writeJson(w, map[string]string{"error": fmt.Sprintf("ошибка при вычислении следующей даты для задачи: %v", err)}, http.StatusInternalServerError)
			return
		}
		err = db.UpdateTaskDate(id, nextDate)
		if err != nil {
			writeJson(w, map[string]string{"error": fmt.Sprintf("ошибка обновления даты задачи: %v", err)}, http.StatusInternalServerError)
			return
		}
		writeJson(w, map[string]int64{}, http.StatusOK)
	}
}
