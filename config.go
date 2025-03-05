package main

import (
	"image"
	"math"
	"time"
)

type config struct {
	img              *image.Image
	angle            float64
	earth            [][]byte
	earth_render     []string
	update_frequency time.Duration
}

func NewConfig(img_fname string, interval time.Duration) *config {

	// updating the world
	ticker := time.NewTicker(interval)

	imgData := read_image(img_fname)

	earth_bytes := make([][]byte, 64)
	for i := range earth_bytes {
		earth_bytes[i] = make([]byte, 2*64)
	}

	earth_render := make([]string, 64)

	cfg := config{img: &imgData,
		angle:            0.0,
		earth:            earth_bytes,
		earth_render:     earth_render,
		update_frequency: interval}

	go func() {
		for {
			<-ticker.C
			cfg.angle += 0.01

			if cfg.angle > 2*math.Pi {
				cfg.angle -= 2 * math.Pi
			}

			// recalculate earth image
			phis, thetas := draw_earth(cfg.angle, *cfg.img, 64)

			cfg.calculate_image(phis, thetas)

			// re-render earth image
			cfg.render_image()
		}
	}()

	return &cfg
}
