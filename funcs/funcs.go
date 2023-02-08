package funcs

import (
	"fmt"
	"math/rand"
)

// Predicate function take a single argument as input and return
// a bool value as output.
func FilterEvenOdd() {

	freq := func(p func(int) bool, vals []int) []int {
		out := make([]int, 0, len(vals))

		for _, v := range vals {
			if p(v) {
				out = append(out, v)
			}
		}

		fmt.Printf("%#v\n", out)
		return out
	}

	isEven := func(n int) bool {
		return n%2 == 0
	}

	isOdd := func(n int) bool {
		return n%2 != 0
	}

	ints := make([]int, 0, 20)
	func() {
		for i := 0; i < 20; i++ {
			ints = append(ints, rand.Intn(20))
		}
	}()

	fmt.Printf("%#v\n", ints)
	freq(isOdd, ints)
	freq(isEven, ints)
}
