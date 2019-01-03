/**
 * Tests for OpenSimplex noise, based on the output of
 * the Java implementation.
 *
 * All reference samples were rendered with the default seed (0). Each version
 * of the noise function (2D, 3D and 4D) was run to output 2D samples slicing
 * across two of the function's axes. There is one 2D slice, three 3D slices
 * and 6 4D slices; the 3D slices each pin one axis to the value 3.8; 4D slices
 * pin one axis (the first in the filename) to 3.8 and the second to 2.7. These
 * values were chosen arbitrarily.
 *
 * Each sample is a 512x512 greyscale PNG; each pixel is 1/24 wide in the
 * OpenSimplex's space -- i.e. pixel (24, 24) in the 2D noise sample was
 * computed by evaluating the 2D noise at (1.0, 1.0) and converting from a [-1,
 * +1] scale to [0, +1].
 */
package ospxn_test

import (
	"testing"

	"github.com/thee-engineer/ospxn"
)

func TestDummy(t *testing.T) {
	noise := ospxn.New(0)

	noise.Eval2(0.0, 0.0)
	noise.Eval3(0.0, 0.0, 0.0)
	noise.Eval4(0.0, 0.0, 0.0, 0.0)
}
