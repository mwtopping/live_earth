package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
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

func show_image(img [][]float32) []string {

	my_strings := make([]string, 0)

	//	shades = " .:-=+*#%@"
	shades = "@%#*+=-:. "
	for _, row := range img {
		this_row := ""
		for _, c := range row {
			index := int(c * 8)
			this_row += string(shades[index])
		}
		my_strings = append(my_strings, this_row)
	}

	return my_strings

}

func show_raw_image(phis, thetas [][]float32, image image.Image) []string {

	my_strings := make([]string, 0)

	shades = " .:-=+*#%@"
	for ip, row := range phis {
		this_row := ""
		for it, _ := range row {
			phi := 512 - int(phis[ip][it])
			theta := int(thetas[ip][it])
			vals := image.At(theta, phi)
			_, _, b, _ := vals.RGBA()
			//				this_row += fmt.Sprintf("%d ", b/8192)
			index := b / 8192

			this_row += string(shades[index])
		}
		my_strings = append(my_strings, this_row)
	}

	return my_strings

}

func sample_image(img [][]float32) []string {

	my_strings := make([]string, 0)

	//	shades = " .:-=+*#%@"
	for _, row := range img {
		this_row := ""
		for i, c := range row {
			if i%2 == 0 {
				this_row += fmt.Sprintf("%.1f ", c)
			}
		}
		my_strings = append(my_strings, this_row)
	}

	return my_strings

}

func reflect_vector(v, norm mat.VecDense) mat.VecDense {
	refl := mat.NewVecDense(3, nil)

	dot := mat.Dot(&v, &norm)

	temp := mat.NewVecDense(norm.Len(), nil)
	temp.ScaleVec(2*dot, &norm)

	refl.SubVec(&v, temp)

	return *refl
}

func normalize(vec *mat.VecDense) mat.VecDense {
	norm := vec.Norm(2)

	normed_vector := mat.NewVecDense(3, nil)

	normed_vector.ScaleVec(1/norm, vec)

	return *normed_vector
}

// draws a simple sphere with phong shading model implemented
func draw_ball(angle float64) []string {

	nx := 41
	ny := 81

	// create blank image
	img := make([][]float32, nx)
	for i := range img {
		img[i] = make([]float32, ny)
	}

	light_pos := mat.NewVecDense(3, []float64{math.Sin(angle), math.Cos(angle), 0.5})

	ones := make([]float64, 3)
	for i := range ones {
		ones[i] = -1
	}
	id := mat.NewDiagDense(3, ones)

	// main loop
	for ix := range nx {
		for iy := range ny {
			x := 0.05 * float64(ix-20)
			y := 0.025 * float64(iy-40)

			// projected distance from sphere
			r := math.Sqrt(x*x + y*y)

			if r < 1.0 {
				sx := x
				sy := y
				sz := math.Sqrt(1 - x*x - y*y)
				sphere_loc := mat.NewVecDense(3, []float64{sx, sy, sz})
				norm_sphere_loc := normalize(sphere_loc)

				vi := mat.NewVecDense(3, nil)
				//				vi.AddScaledVec(&norm_sphere_loc, -1.0, light_pos)
				vi.SubVec(light_pos, mat.NewVecDense(3, []float64{0, 0, 0}))
				*vi = normalize(vi)

				nmatrix := mat.NewDense(norm_sphere_loc.Len(),
					norm_sphere_loc.Len(), nil)

				nmatrix.Outer(1.0, &norm_sphere_loc, &norm_sphere_loc)
				nmatrix.Scale(2.0, nmatrix)

				nmatrix.Add(nmatrix, id)

				si := mat.NewVecDense(3, nil)
				si.MulVec(nmatrix, vi)

				diffuse := 0.15
				direct := clip(0.35 * mat.Dot(vi, &norm_sphere_loc))
				specular := 0.5 * math.Pow(clip(mat.Dot(mat.NewVecDense(3, []float64{0, 0, 1}), si)), 6)

				img[ix][iy] = float32(diffuse + direct + specular)

			}
		}
	}

	return show_image(img)
}

func draw_earth(angle float64, earth_image image.Image) []string {

	//	nx := 41
	//	ny := 81

	nx := 81
	ny := 161

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
			x := 0.025 * float64(ix-40)
			y := 0.0125 * float64(iy-80)

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

	return show_raw_image(phis, thetas, earth_image)

}
