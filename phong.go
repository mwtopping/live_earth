package main

import (
	"gonum.org/v1/gonum/mat"
	"math"
)

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

	return nil
	// return show_image(img)
}
