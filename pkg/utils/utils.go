package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// RandIntInRange creates a new random integer in the specified range.
func RandIntInRange(min, max int) int {
	return int(rand.Int63n(int64(max-min))) + min
}

// Min returns the minimum of the two ints.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// PadOrTrim pads or trims a byte array to the specified size.
func PadOrTrim(bb []byte, size int) []byte {
	l := len(bb)
	if l == size {
		return bb
	}
	if l > size {
		return bb[l-size:]
	}
	tmp := make([]byte, size)
	copy(tmp[size-l:], bb)
	return tmp
}
