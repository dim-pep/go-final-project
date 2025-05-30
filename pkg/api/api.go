package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/task/done", doneHandler)
}

func nextDayHandler(w http.ResponseWriter, r *http.Request) {

	nowStr := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	now, err := time.Parse(dateFormat, nowStr)
	if err != nil {
		http.Error(w, "Ошибка парсинга now", http.StatusBadRequest)
		return
	}

	nextDate, err := NextDate(now, date, repeat)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка вычисления даты: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", nextDate)
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		addTaskHandler(w, r)
	case "GET":
		getTaskHandler(w, r)
	case "PUT":
		updateTaskHandler(w, r)
	case "DELETE":
		deleteTaskHandler(w, r)
	default:
		http.Error(w, "Неизвестный метод", http.StatusMethodNotAllowed)
	}
}

func writeJson(w http.ResponseWriter, data any, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	dataJson, err := json.Marshal(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при сириализации ответа: %v", err), http.StatusInternalServerError)
	}
	w.Write(dataJson)
}
