package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

var shades string

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
			fmt.Printf("%v", string(shades[index]))
			this_row += string(shades[index])
		}
		fmt.Println()
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

func draw_ball(angle float64) []string {

	fmt.Println("Drawing Circle")

	nx := 41
	ny := 81

	// create blank image
	img := make([][]float32, nx)
	for i := range img {
		img[i] = make([]float32, ny)
	}

	light_pos := mat.NewVecDense(3, []float64{math.Sin(angle), math.Cos(angle), 0.5})
	fmt.Println(light_pos)

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
				fmt.Println(si)

				diffuse := 0.15
				direct := clip(0.35 * mat.Dot(vi, &norm_sphere_loc))
				specular := 0.5 * math.Pow(clip(mat.Dot(mat.NewVecDense(3, []float64{0, 0, 1}), si)), 6)

				img[ix][iy] = float32(diffuse + direct + specular)

			}
		}
	}

	return show_image(img)
}

// ReflectVector computes the reflection of vector v across the normal vector norm
// using the formula: refl = v - 2 * dot(v, norm) * norm
// Both vectors must be unit vectors

//def reflect_vector(v, norm):
//
//    norm = normalize(norm)
//    refl = v - 2 * np.dot(v, norm) * norm
//    return refl

//
//shades = " .:-=+*#%@"
//
//
//for ia, ang in tqdm(enumerate(np.linspace(0, 6.28, 100)), total=100):
//
//    str_img = ""
//
//    lightpos = np.array([np.sin(ang), np.cos(ang), 0.2])
//    sloc = np.array([0,0,0])
//
//    for ix, x in enumerate(xs):
//        for iy, y in enumerate(ys):
//            amount = 0
//            ## projected distance from center of sphere
//            r = np.sqrt(x*x+y*y)
//            if r <= 1: # in sphere
//                sx = x
//                sy = y
//                sz = np.sqrt(1-x*x-y*y)
//
//                norm = normalize(np.array([sx, sy, sz]))
//                vi = normalize(lightpos-sloc)
//
//                nmatrix = norm.copy().reshape(1, 3)
//
//    #            si = (2*np.matmul(np.matrix(norm), np.matrix(norm).T))
//                si = np.matmul(2*np.outer( norm, norm) - np.identity(3), vi)
//                vr = reflect_vector(vi, norm)
//
//
//
//                diffuse = 0.15
//
//                direct = 0.35*np.dot(vi,norm)
//
//                specular = 0.5*(np.max((0, np.dot(np.array([0,0,1]), si)))**6)
//
//                amount = diffuse + np.max((0, direct)) + specular
//
//                img[ix][iy] = diffuse + np.max((0, direct)) + specular
//
//#            str_img += shades[int(amount*70)]
//            str_img += shades[int(amount*nshades)]
//        str_img += "\n"
//
//
//
//#    fig, ax = plt.subplots()
//
//
//    print(r"{}".format(str_img))
//    print(len(str_img))
//#    ax.imshow(img, origin='lower')
//#    plt.show()
//    plt.close('all')
