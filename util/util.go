package util

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomNumberRange
func RandomNumberRange() {
	IntRange := func(min, max int) int {
		return rand.Intn(max-min) + min
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		fmt.Printf("Value %v: %v\n", i, IntRange(10, 20))
	}
}
