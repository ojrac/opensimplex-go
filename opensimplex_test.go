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
package opensimplex

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"os"
	"path"
	"testing"
)

func loadSamples() <-chan []float64 {
	c := make(chan []float64)
	go func() {
		f, err := os.Open(path.Join("test_files", "samples.json.gz"))
		if err != nil {
			panic(err.Error())
		}
		defer f.Close()

		gz, err := gzip.NewReader(f)
		if err != nil {
			panic(err.Error())
		}

		dec := json.NewDecoder(gz)
		for {
			var sample []float64
			if err := dec.Decode(&sample); err == io.EOF {
				break
			} else if err != nil {
				panic(err.Error())
			} else {
				c <- sample
			}
		}
		close(c)
	}()

	return c
}

func TestSamplesMatch(t *testing.T) {
	samples := loadSamples()
	n := New(0)

	for s := range samples {
		var expected, actual float64
		switch len(s) {
		case 3:
			expected = s[2]
			actual = n.Eval2(s[0], s[1])
		case 4:
			expected = s[3]
			actual = n.Eval3(s[0], s[1], s[2])
		case 5:
			expected = s[4]
			actual = n.Eval4(s[0], s[1], s[2], s[3])
		default:
			t.Fatalf("Unexpected size sample: %d", len(s))
		}

		if expected != actual {
			t.Fatalf("Expected %v, got %v for %dD sample at %v",
				expected, actual, len(s)-1, s[:len(s)-1])
		}
	}
}
