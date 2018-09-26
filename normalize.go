package opensimplex

const (
	// The normMin and normScale constants are used
	// in the formula for normalizing the raw output
	// of the OpenSimplex algorithm. They were
	// derived from empirical observations of the
	// range of raw values.
	normMin   = 0.864366441
	normScale = 0.5784583670
)

// Norm accepts a value from one of the Eval functions and
// normalizes it to a value in the range [0, 1). Note: if
// the desired end result is a float32, use
// Norm32 instead, which correctly handles the reduction
// in precision. If the desired end result is an int, use
// NormInt.
func Norm(n float64) float64 {
	return (n + normMin) * normScale
}

// Norm32 accepts a value from one of the Eval functions and
// normalizes it to a value in the range [0, 1). Because a
// simple cast from a float64 to a float32 can cause issues
// with precision, this function should be used instead of a
// cast whenever a float32 is desired.
func Norm32(n float64) float32 {
	norm64 := (n + normMin) * normScale
	norm32 := float32(norm64)

	if norm32 >= 1.0 {
		return float32(0.999999)
	} else {
		return norm32
	}
}

// NormInt accepts a value from one of the Eval functions
// and normalizes it to an int in the range [0, maxExclusive).
// Because a simple cast from a float64 to an int can cause
// issues with precision, this function should be used instead
// of a cast whenever an int is desired.
func NormInt(n float64, maxExclusive int) int {
	norm64 := (n + normMin) * normScale
	i := int(norm64 * float64(maxExclusive))

	if i == maxExclusive {
		return maxExclusive - 1
	} else {
		return i
	}
}
