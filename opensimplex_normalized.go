package opensimplex

const (
	// The normMin and normScale constants are used
	// in the formula for normalizing the raw output
	// of the OpenSimplex algorithm. They were
	// derived from empirical observations of the
	// range of raw values. Different constants are
	// required for each of Eval2, Eval3, and Eval4.
	normMin2   = 0.864366441
	normScale2 = 0.5784583670

	normMin3   = 0.944824004155
	normScale3 = 0.5291990849667171

	normMin4   = 1.034
	normScale4 = 0.48355899419729204
)

type normNoise struct {
	base Noise
}

// Eval2 returns a random noise value in two dimensions
// in the range [0, 1).
func (s *normNoise) Eval2(x, y float64) float64 {
	//return norm2_64(s.base.Eval2(x, y))
	r := s.base.Eval2(x, y)
	return (r + normMin2) * normScale2
}

// Eval3 returns a random noise value in three dimensions
// in the range [0, 1).
func (s *normNoise) Eval3(x, y, z float64) float64 {
	r := s.base.Eval3(x, y, z)
	return (r + normMin3) * normScale3
}

// Eval4 returns a random noise value in four dimensions
// in the range [0, 1).
func (s *normNoise) Eval4(x, y, z, t float64) float64 {
	r := s.base.Eval4(x, y, z, t)
	return (r + normMin4) * normScale4
}

type normNoise32 struct {
	base Noise
}

// Eval2 returns a random noise value in two dimensions
// in the range [0, 1).
func (s *normNoise32) Eval2(x, y float32) float32 {
	r := s.base.Eval2(float64(x), float64(y))
	norm64 := (r + normMin2) * normScale2
	norm32 := float32(norm64)

	// Empirical testing shows that a simple float32 cast
	// from the normalized float64, as above, will sometimes
	// produce a value of 1.0.
	if norm32 >= 1.0 {
		return float32(0.999999)
	} else {
		return norm32
	}
}

// Eval3 returns a random noise value in three dimensions
// in the range [0, 1).
func (s *normNoise32) Eval3(x, y, z float32) float32 {
	r := s.base.Eval3(float64(x), float64(y), float64(z))
	norm64 := (r + normMin3) * normScale3
	norm32 := float32(norm64)

	// Unlike Eval2, have not actually tested whether a
	// simple float32 cast will produce 1.0, but it seems likely.
	if norm32 >= 1.0 {
		return float32(0.999999)
	} else {
		return norm32
	}
}

// Eval4 returns a random noise value in four dimensions
// in the range [0, 1).
func (s *normNoise32) Eval4(x, y, z, t float32) float32 {
	r := s.base.Eval4(float64(x), float64(y), float64(z), float64(t))
	norm64 := (r + normMin4) * normScale4
	norm32 := float32(norm64)

	// Unlike Eval2, have not actually tested whether a
	// simple float32 cast will produce 1.0, but it seems likely.
	if norm32 >= 1.0 {
		return float32(0.999999)
	} else {
		return norm32
	}
}
