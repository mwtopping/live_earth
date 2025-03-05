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

	// special behavior for mobile platform
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
				for _, row := range cfg.earth_render {
					fmt.Fprintf(w, "data: %v\n", row)
				}

			case "mobile":
				// make mobile version a factor of 4 smaller
				for i := 0; i < len(cfg.earth_render); i += 4 {
					fmt.Fprintf(w, "data: %v\n", extract_periodic(cfg.earth_render[i], 4))
				}

			}
			fmt.Fprintf(w, "\n")
			flusher.Flush()

			time.Sleep(cfg.update_frequency)

		}
	}
}

func extract_periodic(s string, n int) string {
	b := make([]byte, 0)

	//for _, c := range string {
	for i := 0; i < len(s); i += n {
		b = append(b, s[i])
	}

	return string(b)
}
