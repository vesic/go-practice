package complex

import (
	"fmt"
	"math/rand"
	"sort"
)

// Perform sorting on slices using the sort package.
// By default, the slice will be sorted from lowest to highest.
func sortingSlices() {
	numbers := []int{}

	for i := 0; i < 10; i++ {
		numbers = append(numbers, rand.Intn(100))
	}

	fmt.Printf("%#v\n", numbers)

	sort.Ints(numbers)

	fmt.Printf("%#v\n", numbers)
}
