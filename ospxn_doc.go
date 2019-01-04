// Package ospxn is a Go implementation of Kurt Spencer's patent-free alternative
// to Perlin and Simplex noise, with minor changes.
//
// Given a seed, it generates smoothly-changing deterministic random values in
// 2 or 3 dimensions. It's commonly used for procedurally generated images,
// geometry, or other randomly-influenced applications that require a random
// gradient.
//
// For more information on OpenSimplex noise, read more from the creator of the
// algorithm: http://uniblock.tumblr.com/post/97868843242/noise
package ospxn
