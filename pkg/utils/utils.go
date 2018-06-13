package utils

import (
	"math/rand"
	"time"
)

// RandIntInRange creates a new random integer in the specified range.
func RandIntInRange(min, max int32) int32 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Int31n(max-min) + min
}
