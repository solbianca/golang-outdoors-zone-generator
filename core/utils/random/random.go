package random

import (
	"fmt"
	"math/rand"
	"time"
)

var rg *randomGenerator

func init() {
	seed := rand.NewSource(time.Now().UnixNano()).Int63()
	rg = newRandomGenerator(seed)
}

type randomGenerator struct {
	seed      int64
	generator *rand.Rand
}

func newRandomGenerator(seed int64) *randomGenerator {
	return &randomGenerator{seed: seed, generator: rand.New(rand.NewSource(seed))}
}

func Between(min, max int) int {
	if min > max {
		panic(fmt.Sprintf("Min value cannot be greater than max. Min:[%d] Max:[%d].", min, max))
	}

	return rg.generator.Intn(max-min+1) + min
}
