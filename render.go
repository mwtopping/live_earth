package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

func show_image(img [][]float32) {
	for _, row := range img {
		fmt.Println(row)
	}
}

func draw_ball() {

	fmt.Println("Drawing Circle")

	nx := 21
	ny := 21

	// create blank image
	img := make([][]float32, nx)
	for i := range img {
		img[i] = make([]float32, ny)
	}

	light_pos := mat.NewVecDense(3, []float64{1.0, 1.0, 0.2})
	fmt.Println(light_pos)

	// main loop
	for ix := range nx {
		for iy := range ny {
			x := 0.1 * float64(ix-10)
			y := 0.1 * float64(iy-10)

			// projected distance from sphere
			r := math.Sqrt(x*x + y*y)

			if r < 1.0 {
				img[ix][iy] = 1
			}
		}
	}

	show_image(img)
}

//def normalize(vec):
//    return vec/np.linalg.norm(vec)

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
