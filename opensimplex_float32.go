package opensimplex

// Wraps a noise instance to work with float32 values

type cast32Noise struct {
	base Noise
}

func (n *cast32Noise) Eval2(x, y float32) float32 {
	return float32(n.base.Eval2(float64(x), float64(y)))
}
func (n *cast32Noise) Eval3(x, y, z float32) float32 {
	return float32(n.base.Eval3(float64(x), float64(y), float64(z)))
}
func (n *cast32Noise) Eval4(x, y, z, w float32) float32 {
	return float32(n.base.Eval4(float64(x), float64(y), float64(z), float64(w)))
}
