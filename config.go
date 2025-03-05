package main

import (
	"image"
	"math"
	"time"
)

type config struct {
	img   *image.Image
	angle float64
}

func NewConfig(img_fname string, interval time.Duration) *config {

	// updating the world
	ticker := time.NewTicker(interval)

	imgData := read_image(img_fname)

	cfg := config{img: &imgData, angle: 0.0}

	go func() {
		for {
			<-ticker.C
			cfg.angle += 0.01

			if cfg.angle > 2*math.Pi {
				cfg.angle -= 2 * math.Pi
			}
		}
	}()

	return &cfg
}
