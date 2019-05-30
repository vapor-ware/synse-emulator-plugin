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

// Min returns the minimum of the two integers.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
