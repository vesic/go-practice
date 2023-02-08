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

// Program to count how many times a word appears in the text.
// Common task in text processing know as word frequency.
func wordFrequency() {
	frequency := func(words []string) map[string]int {
		freq := make(map[string]int)
		for _, w := range words {
			freq[w]++
		}
		return freq
	}

	sorted := frequency([]string{
		"be", "and", "of", "to", "have", "be", "and", "of",
	})
	fmt.Printf("%#v\n", sorted)
}
