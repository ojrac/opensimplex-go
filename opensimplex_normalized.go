package opensimplex

const (
	// The normMin and normScale constants are used
	// in the formula for normalizing the raw output
	// of the OpenSimplex algorithm. They were
	// derived from empirical observations of the
	// range of raw values. They work on the approximate
	// (-0.866, 0.866) range produced by Eval2.
	// TODO: If the raw Eval3 or Eval4 functions produce a
	// different range, they will require different constants to
	// normalize.
	normMin   = 0.864366441
	normScale = 0.5784583670
)

// When Eval3 and Eval4 methods are added, this will
// implement the Noise interface.
type noiseNormalized struct {
	noise
}

// Eval2 returns a random noise value in two dimensions
// in the range [0, 1).
func (s *noiseNormalized) Eval2(x, y float64) float64 {
	return norm(s.noise.Eval2(x, y))
}

// When Eval3 and Eval4 methods are added, this will
// implement the Noise32 interface.
type noiseNormalized32 struct {
	noise
}

// Eval2 returns a random noise value in two dimensions
// in the range [0, 1).
func (s *noiseNormalized32) Eval2(x, y float32) float32 {
	return norm32(s.noise.Eval2(float64(x), float64(y)))
}

type noiseNormalizedInt struct {
	noise
}

// Eval2 returns a random noise value in two dimensions
// in the range [0, maxExclusive).
func (s *noiseNormalizedInt) Eval2(x, y, maxExclusive int) int {
	return normInt(s.noise.Eval2(float64(x), float64(y)), maxExclusive)
}

// norm accepts a value from one of the float64 Eval functions
// and normalizes it to a value in the range [0, 1). Note: if
// the desired end result is a float32, use norm32 instead,
// which correctly handles the reduction in precision. If the
// desired end result is an int, use normInt.
func norm(n float64) float64 {
	return (n + normMin) * normScale
}

// norm32 accepts a value from one of the float64 Eval functions and
// normalizes it to a value in the range [0, 1). Because a
// simple cast from a float64 to a float32 can cause issues
// with precision, this function should be used instead of a
// cast whenever a float32 is desired.
func norm32(n float64) float32 {
	norm64 := (n + normMin) * normScale
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

// normInt accepts a value from one of the Eval functions
// and normalizes it to an int in the range [0, maxExclusive).
// Because a simple cast from a float64 to an int can cause
// issues with precision, this function should be used instead
// of a cast whenever an int is desired.
func normInt(n float64, maxExclusive int) int {
	norm64 := (n + normMin) * normScale
	i := int(norm64 * float64(maxExclusive))

	if i == maxExclusive {
		return maxExclusive - 1
	} else {
		return i
	}
}
