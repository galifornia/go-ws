package main

import (
	"log"
	"net/http"

	"github.com/galifornia/go-ws/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting go routine that broadcasts messages")
	go handlers.ListenWsChannel()

	log.Println("Starting web server on port 8080")

	_ = http.ListenAndServe("localhost:8080", mux)
}
