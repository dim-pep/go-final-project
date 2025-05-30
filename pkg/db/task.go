package db

import "fmt"

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`
	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func Tasks(limit int) ([]*Task, error) {
	var tasks []*Task
	query := `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date DESC LIMIT ?`
	rows, err := db.Query(query, limit)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса SELECT: %v", err)
	}
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, fmt.Errorf("ошибка при чтении данных из базы данных: %v", err)
		}
		tasks = append(tasks, &task)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при чтении данных из базы данных: %v", err)
	}
	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`
	row := db.QueryRow(query, id)
	var task Task
	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении данных из базы данных: %v", err)
	}
	return &task, nil
}
