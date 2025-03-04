package random

import (
	"math/rand"
	"time"
)

// Number returns a random number of type T between min and max (inclusive).
func Number[T int | int64 | float64](min, max T) T {
	if min > max {
		panic("min cannot be greater than max")
	}
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)

	switch any(min).(type) {
	case int:
		return T(r.Intn(int(max-min+1)) + int(min))
	case int64:
		return T(r.Int63n(int64(max-min+1)) + int64(min))
	case float64:
		return T(min + T(r.Float64())*(max-min))
	default:
		panic("unsupported type")
	}
}
