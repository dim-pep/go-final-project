package server

import (
	"fmt"
	"net/http"

	"go1f/pkg/api"
)

func Start() error {
	port := 7540
	http.Handle("/", http.FileServer(http.Dir("web")))
	api.Init()
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
