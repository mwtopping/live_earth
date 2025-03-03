package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handleSSE(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Unable to stream", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		select {
		case <-r.Context().Done():
			return
		default:
			//			fmt.Fprintf(w, "data: ==================\n")
			//			fmt.Fprintf(w, "data: ==================\n%v\n=================\n\n", time.Now())
			//			fmt.Fprintf(w, "==================\n\n")
			for i := 0; i < 20; i++ {
				fmt.Fprintf(w, "data: Line %d of event %v\n", i, time.Now())
			}
			fmt.Fprintf(w, "\n")
			flusher.Flush()
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	const port = ":8080"

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./")))
	mux.HandleFunc("/events", handleSSE)

	server := &http.Server{Handler: mux, Addr: port}

	log.Printf("Server started on %v\n", port)
	log.Fatal(server.ListenAndServe())
}
