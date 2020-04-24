package utils

import (
	"math/rand"
)

// RandIntInRange creates a new random integer within the specified range.
func RandIntInRange(min, max int) int {
	return int(rand.Int63n(int64(max-min))) + min
}

// BoundedIncrement increments the starting value within the given bounds.
func BoundedIncrement(start interface{}, lower, upper int) int {
	if start == nil {
		return lower
	}
	s := start.(int)
	s++

	if s > upper {
		return lower
	}
	return s
}

// RandWalkInRange creates a new value by walking a random distance from the
// start value.
func RandWalkInRange(start interface{}, lower, upper int) int {
	// Between readings, we allow the value to change between 0 and 4.
	// This allows for some movement, but doesn't allow massive swings in
	// short periods.
	diff := RandIntInRange(0, 4)

	// If a seed is not set, start at the midway point between the lower and upper bounds.
	if start == nil {
		start = int((lower + upper) / 2)
	}

	// Determine whether the value will be a positive or negative step
	positive := rand.Intn(2) == 0

	if positive {
		newVal := start.(int) + diff
		if newVal > upper {
			return start.(int) - diff
		}
		return newVal
	}

	newVal := start.(int) - diff
	if newVal < lower {
		return start.(int) + diff
	}
	return newVal
}
