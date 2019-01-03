package ospxn

const (
	stretchConstant2D float64 = -0.211324865405187 // (1/Math.sqrt(2+1)-1)/2
	squishConstant2D  float64 = 0.366025403784439  // (Math.sqrt(2+1)-1)/2
	normConstant2D    float64 = 47

	stretchConstant3D float64 = -1.0 / 6 // (1/Math.sqrt(3+1)-1)/3
	squishConstant3D  float64 = 1.0 / 3  // (Math.sqrt(3+1)-1)/3
	normConstant3D    float64 = 103
)

func (n *noise) extrapolate2(xsb, ysb int32, dx, dy float64) float64 {
	index := n.perm[(int32(n.perm[xsb&0xFF])+ysb)&0xFF] & 0x0E
	return float64(gradients2D[index])*dx + float64(gradients2D[index+1])*dy
}

func (n *noise) extrapolate3(xsb, ysb, zsb int32, dx, dy, dz float64) float64 {
	index := n.perm3[(int32(n.perm[(int32(n.perm[xsb&0xFF])+ysb)&0xFF])+zsb)&0xFF]
	return float64(gradients3D[index])*dx + float64(gradients3D[index+1])*dy + float64(gradients3D[index+2])*dz
}

// Gradients for 2D. They approximate the directions to the
// vertices of an octagon from the center.
var gradients2D = []int8{
	5, 2, 2, 5,
	-5, 2, -2, 5,
	5, -2, 2, -5,
	-5, -2, -2, -5,
}

// Gradients for 3D. They approximate the directions to the
// vertices of a rhombicuboctahedron from the center, skewed so
// that the triangular and square facets can be inscribed inside
// circles of the same radius.
var gradients3D = []int8{
	-11, 4, 4, -4, 11, 4, -4, 4, 11,
	11, 4, 4, 4, 11, 4, 4, 4, 11,
	-11, -4, 4, -4, -11, 4, -4, -4, 11,
	11, -4, 4, 4, -11, 4, 4, -4, 11,
	-11, 4, -4, -4, 11, -4, -4, 4, -11,
	11, 4, -4, 4, 11, -4, 4, 4, -11,
	-11, -4, -4, -4, -11, -4, -4, -4, -11,
	11, -4, -4, 4, -11, -4, 4, -4, -11,
}
