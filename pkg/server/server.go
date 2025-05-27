package server

import (
	"fmt"
	"net/http"
)

func Start() error {
	port := 7540
	http.Handle("/", http.FileServer(http.Dir("web")))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
