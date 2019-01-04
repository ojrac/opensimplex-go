package opensimplex

import "math/rand"

func Example() {
	noise := New(rand.Int63())

	w, h := 100, 100
	heightmap := make([]float64, w*h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			xFloat := float64(x) / float64(w)
			yFloat := float64(y) / float64(h)
			heightmap[(y*w)+x] = noise.Eval2(xFloat, yFloat)
		}
	}
}
