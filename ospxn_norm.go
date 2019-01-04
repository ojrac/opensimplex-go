package ospxn

const (
	normMin2   float64 = 0.864366441
	normScale2 float64 = 0.5784583670
	normMin3   float64 = 0.944824004155
	normScale3 float64 = 0.5291990849667171
)

type normNoise struct {
	n Noise
}

// NewNormalized constructs a normalized Noise instance with a 64-bit seed.
// Eval methods will return values in [0, 1).
func NewNormalized(seed int64) Noise {
	return &normNoise{n: New(seed)}
}

// Eval2 returns a random noise value in two dimensions in the range [0, 1).
func (nn *normNoise) Eval2(x, y float64) float64 {
	r := nn.n.Eval2(x, y)
	return (r + normMin2) * normScale2
}

// Eval3 returns a random noise value in three dimensions in the range [0, 1).
func (nn *normNoise) Eval3(x, y, z float64) float64 {
	r := nn.n.Eval3(x, y, z)
	return (r + normMin3) * normScale3
}
