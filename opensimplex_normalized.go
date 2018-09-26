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
)

// TODO: Implement Eval3 and Eval4 to satisfy Noise interface.
type normNoise struct {
	base noise
}

// Eval2 returns a random noise value in two dimensions
// in the range [0, 1).
func (s *normNoise) Eval2(x, y float64) float64 {
	//return norm2_64(s.base.Eval2(x, y))
	r := s.base.Eval2(x, y)
	return (r + normMin2) * normScale2
}

func (s *normNoise) Eval3(x, y, z float64) float64 {
	r := s.base.Eval3(x, y, z)
	return (r + normMin3) * normScale3
}

// TODO: Implement Eval3 and Eval4 to satisfy Noise32 interface.
type normNoise32 struct {
	base noise
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
