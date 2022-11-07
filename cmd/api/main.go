package main

import (
	"net/http"

	httpadapter "github.com/volvofixthis/repository-badge/internal/adapter/http"
)

func main() {
	router := httpadapter.NewRouter()
	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	srv.ListenAndServe()

}
