package main

import (
	"adias-alcove/handlers"
	"net/http"
)

func registerHandlers() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.Handle("/static/", handlers.StaticHandler())
}

func main() {
	registerHandlers()
	http.ListenAndServe(":9090", nil)
}
