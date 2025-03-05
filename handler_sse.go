package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func (cfg *config) handleSSE(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Unable to stream", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	log.Println(r.UserAgent())

	device := "desktop"
	if strings.Contains(r.UserAgent(), "Mobile") {
		device = "mobile"
	}

	for {
		select {
		case <-r.Context().Done():
			log.Println("Connection Closed")
			return
		default:
			switch device {
			case "desktop":
				image := draw_earth(cfg.angle, *cfg.img, 64)

				for _, row := range image {
					fmt.Fprintf(w, "data: %v\n", row)
				}

			case "mobile":
				image := draw_earth(cfg.angle, *cfg.img, 16)

				for _, row := range image {
					fmt.Fprintf(w, "data: %v\n", row)
				}

			}
			fmt.Fprintf(w, "\n")
			flusher.Flush()

			// probably want this to be the same as the angle update frequency
			time.Sleep(100 * time.Millisecond)

		}
	}
}
