package random

import (
	"math/rand"
	"time"
)

// Bool returns a random boolean value (true or false).
func Bool() bool {
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)
	return r.Intn(2) == 1
}
