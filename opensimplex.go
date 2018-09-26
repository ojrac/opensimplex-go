// opensimplex is a Go implementation of Kurt Spencer's patent-free alternative
// to Perlin and Simplex noise.
//
// Given a seed, it generates smoothly-changing deterministic random values in
// 2, 3 or 4 dimensions. It's commonly used for procedurally generated images,
// geometry, or other randomly-influenced applications that require a random
// gradient.
//
// For more information on OpenSimplex noise, read more from the creator of the
// algorithm: http://uniblock.tumblr.com/post/97868843242/noise
package opensimplex

/**
 * OpenSimplex Noise in Go.
 * algorithm by Kurt Spencer
 * ported by Owen Raccuglia
 *
 * Based on Java v1.1 (October 5, 2014)
 */

// A seeded 64-bit noise instance
type Noise interface {
	Eval2(x, y float64) float64
	Eval3(x, y, z float64) float64
	Eval4(x, y, z, w float64) float64
}

// A seeded 32-bit noise instance
type Noise32 interface {
	Eval2(x, y float32) float32
	Eval3(x, y, z float32) float32
	Eval4(x, y, z, w float32) float32
}

// Returns a Noise instance with a seed of 0.
func New() Noise {
	return NewWithSeed(defaultSeed)
}

// Returns a Noise instance with a 64-bit seed. Two Noise instances with the
// same seed will have the same output.
func NewWithSeed(seed int64) Noise {
	s := &noise{
		perm:            make([]int16, 256),
		permGradIndex3D: make([]int16, 256),
	}

	source := make([]int16, 256)
	for i := range source {
		source[i] = int16(i)
	}

	seed = seed*6364136223846793005 + 1442695040888963407
	seed = seed*6364136223846793005 + 1442695040888963407
	seed = seed*6364136223846793005 + 1442695040888963407
	for i := int32(255); i >= 0; i-- {
		seed = seed*6364136223846793005 + 1442695040888963407
		r := int32((seed + 31) % int64(i+1))
		if r < 0 {
			r += i + 1
		}

		s.perm[i] = source[r]
		s.permGradIndex3D[i] = (s.perm[i] % (int16(len(gradients3D)) / 3)) * 3
		source[r] = source[i]
	}

	return s
}

// Returns a Noise instance with a specific internal permutation state.
// If you're not sure about this, you probably want NewWithSeed().
func NewWithPerm(perm []int16) Noise {
	s := &noise{
		perm:            perm,
		permGradIndex3D: make([]int16, 256),
	}

	for i, p := range perm {
		// Since 3D has 24 gradients, simple bitmask won't work, so precompute modulo array.
		s.permGradIndex3D[i] = (p % (int16(len(gradients3D)) / 3)) % 3
	}

	return s
}

// Wraps a noise instance to work with float32 values
func AsNoise32(noise Noise) Noise32 {
	return &noiseCast32{noise}
}

type noiseCast32 struct {
	noise Noise
}

func (n *noiseCast32) Eval2(x, y float32) float32 {
	return float32(n.noise.Eval2(float64(x), float64(y)))
}
func (n *noiseCast32) Eval3(x, y, z float32) float32 {
	return float32(n.noise.Eval3(float64(x), float64(y), float64(z)))
}
func (n *noiseCast32) Eval4(x, y, z, w float32) float32 {
	return float32(n.noise.Eval4(float64(x), float64(y), float64(z), float64(w)))
}
