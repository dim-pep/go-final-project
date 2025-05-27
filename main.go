package main

import (
	"go1f/pkg/db"
	"go1f/pkg/server"
	"log"
)

func main() {
	dbFile := "scheduler.db"

	if err := db.Init(dbFile); err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}

	if err := server.Start(); err != nil {
		log.Fatalf("Ошибка запуска веб-сервера: %v", err)
	}
}
