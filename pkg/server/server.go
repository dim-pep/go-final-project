package server

import (
	"fmt"
	"log"
	"net/http"

	"go1f/pkg/api"
)

func Start() error {
	port := 7540
	log.Printf("Приложение запущено на порту %d", port)
	http.Handle("/", http.FileServer(http.Dir("web")))
	api.Init()
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
