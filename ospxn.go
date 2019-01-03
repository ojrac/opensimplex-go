package ospxn

// Noise is the interface for the noise generator (2, 3 and 4 D)
type Noise interface {
	Eval2(x, y float64) float64
	Eval3(x, y, z float64) float64
}

// New constructs a Noise instance with a 64-bit seed.
func New(seed int64) Noise {
	n := &noise{
		perm:  make([]int16, 256),
		perm3: make([]int16, 256),
	}

	var idx int32

	source := make([]int16, 256)
	for idx := range source {
		source[idx] = int16(idx)
	}

	seed = seed*6364136223846793005 + 1442695040888963407
	seed = seed*6364136223846793005 + 1442695040888963407
	seed = seed*6364136223846793005 + 1442695040888963407

	for idx = 255; idx >= 0; idx-- {
		seed = seed*6364136223846793005 + 1442695040888963407
		r := int32((seed + 31) % int64(idx+1))
		if r < 0 {
			r += idx + 1
		}

		n.perm[idx] = source[r]
		n.perm3[idx] = (n.perm[idx] % (int16(len(gradients3D)) / 3)) * 3
		source[r] = source[idx]
	}

	return n
}
