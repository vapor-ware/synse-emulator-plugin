package utils

import (
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// EmitMode specifies the behavior of a ValueEmitter.
type EmitMode int

const (
	// Store the value on write and return it on read.
	Store EmitMode = iota

	// RandomRange generates a random value within a specified range.
	RandomRange

	// Accumulate increases the starting value for every read.
	Accumulate

	// RandomWalk generates a random value close to its previous value.
	RandomWalk
)

// ValueEmitter provides a way to get emulated reading values. It offers
// multiple behavior modes, so readings from emulated devices may behave
// differently, but be managed by the same interface.
type ValueEmitter struct {
	mode       EmitMode
	mux        sync.Mutex
	prev       interface{}
	upperBound int
	lowerBound int
}

// NewValueEmitter creates a new ValueEmitter for the specified mode.
//
// A starting value may be seeded via WithSeed, and bounds may be set
// via WithUpperBound and WithLowerBound. Those methods return a pointer
// to the emitter, so they can be chained on creation, e.g.
//
//     emitter := NewValueEmitter(RandomWalk).WithSeed(23).WithUpperBound(4).WithLowerBound(1)
func NewValueEmitter(mode EmitMode) *ValueEmitter {
	var (
		upper int
		lower int
	)
	switch mode {
	case RandomRange:
		// Defaults for random range
		upper = 10
		lower = 0
	case Accumulate:
		// Defaults for accumulate
		upper = 1000
		lower = 0
	case RandomWalk:
		// Defaults for random walk
		upper = 8
		lower = 0
	}

	return &ValueEmitter{
		mode:       mode,
		upperBound: upper,
		lowerBound: lower,
	}
}

// Next gets the next value from the ValueEmitter.
func (emitter *ValueEmitter) Next() interface{} {
	emitter.mux.Lock()
	defer emitter.mux.Unlock()

	switch emitter.mode {
	case RandomRange:
		emitter.prev = RandIntInRange(emitter.lowerBound, emitter.upperBound)

	case Accumulate:
		emitter.prev = BoundedIncrement(emitter.prev, emitter.lowerBound, emitter.upperBound)

	case RandomWalk:
		emitter.prev = RandWalkInRange(emitter.prev, emitter.lowerBound, emitter.upperBound)

	case Store:
	}
	return emitter.prev
}

// Set the value for the ValueEmitter.
//
// Generally, this should only be used with Store mode.
func (emitter *ValueEmitter) Set(value interface{}) {
	emitter.mux.Lock()
	defer emitter.mux.Unlock()

	emitter.prev = value
}

// WithSeed is a special case of Set, where it will set the starting emitter
// value. If the emitter already has a value set, this will do nothing.
func (emitter *ValueEmitter) WithSeed(seed interface{}) *ValueEmitter {
	if emitter.prev == nil {
		emitter.prev = seed
	}
	return emitter
}

// WithUpperBound sets the upper bound for the emitter. The upper bound
// behaves differently depending on the run mode.
//
// * Store: No effect.
// * RandomRange: Sets the maximum possible random value.
// * Accumulate: Sets maximum possible value for accumulation, after which the
//   	values are reset.
// * RandomWalk: Sets the maximum possible value to walk to. Once it reaches this
//      bound, it will start forcing the walk downwards.
func (emitter *ValueEmitter) WithUpperBound(bound int) *ValueEmitter {
	emitter.upperBound = bound
	return emitter
}

// WithLowerBound sets the lower bound for the emitter. The lower bound
// behaves differently depending on the run mode.
//
// This does not verify that the lower bound is actually lower than the
// upper bound. It is left to the caller to ensure this is the case.
//
// * Store: No effect.
// * RandomRange: Sets the minimum possible random value.
// * Accumulate: Sets the value which the accumulator will reset to. If no
//   	value is seeded (via Set), this will also be the starting seed.
// * RandomWalk: Sets the minimum possible value to walk to. Once it reaches this
//   	bound, it will start forcing the walk upwards.
func (emitter *ValueEmitter) WithLowerBound(bound int) *ValueEmitter {
	emitter.lowerBound = bound
	return emitter
}
