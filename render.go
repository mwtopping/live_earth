package main

import (
	"image"
	"math"
)

var shades string

func mapRange(value, fromMin, fromMax, toMin, toMax float64) float64 {
	return (value-fromMin)/(fromMax-fromMin)*(toMax-toMin) + toMin
}

func clip(x float64) float64 {
	if x < 0 {
		return 0.0
	}
	return x
}

func (cfg *config) calculate_image(phis, thetas [][]float32) {

	shades = " .:-=+*#%@"
	for ip, row := range phis {
		for it, _ := range row {

			phi := 512 - int(phis[ip][it])
			theta := int(thetas[ip][it])
			vals := (*cfg.img).At(theta, phi)
			_, _, b, _ := vals.RGBA()
			//				this_row += fmt.Sprintf("%d ", b/8192)
			index := b / 8192

			cfg.earth[ip][it] = shades[index]
		}
	}

}

func (cfg *config) render_image() {
	for ir, row := range cfg.earth {
		this_row := string(row)
		cfg.earth_render[ir] = this_row
	}
}

func draw_earth(angle float64, earth_image image.Image, size int) ([][]float32, [][]float32) {

	nx := size
	ny := 2 * size

	// create blank image
	phis := make([][]float32, nx)
	thetas := make([][]float32, nx)
	for i := range phis {
		phis[i] = make([]float32, ny)
		thetas[i] = make([]float32, ny)
	}

	// main loop
	for ix := range nx {
		for iy := range ny {
			x := float64(ix-size/2) / float64(size/2)
			y := float64(iy-size) / float64(size)

			// projected distance from sphere
			r := math.Sqrt(x*x + y*y)

			if r < 1.0 {
				sx := x
				sy := y
				sz := math.Sqrt(1 - x*x - y*y)

				theta := (math.Atan2(sy, sz) + math.Pi) + angle
				if theta > 2*math.Pi {
					theta -= 2 * math.Pi
				}
				phi := math.Acos(sx)

				thetas[ix][iy] = float32(mapRange(theta, 0, 2*math.Pi, 0, 1024))
				phis[ix][iy] = float32(mapRange(phi, 0, math.Pi, 0, 512))
			}
		}
	}

	return phis, thetas

}
