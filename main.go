package main

import (
	"fmt"
	"log"
	"math"
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

	angle := 0.0

	earth_image := read_image("./assets/earth.png")

	for {
		select {
		case <-r.Context().Done():
			return
		default:
			//			image := draw_ball(angle)
			image := draw_earth(angle, earth_image)
			for _, row := range image {
				fmt.Fprintf(w, "data: %v\n", row)
			}
			fmt.Fprintf(w, "\n")
			flusher.Flush()
			angle += 0.05
			if angle > 2*math.Pi {
				angle -= 2 * math.Pi
			}
			time.Sleep(100 * time.Millisecond)
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
