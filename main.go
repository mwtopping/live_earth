package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	const port = ":8080"

	mux := http.NewServeMux()

	cfg := NewConfig("./assets/earth.png", 100*time.Millisecond)

	mux.Handle("/", http.FileServer(http.Dir("./")))
	mux.HandleFunc("/events", cfg.handleSSE)

	server := &http.Server{Handler: mux, Addr: port}

	log.Printf("Server started on %v\n", port)
	log.Fatal(server.ListenAndServe())
}
