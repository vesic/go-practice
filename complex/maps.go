package complex

import (
	"fmt"
	"sort"
)

// EnumeratingMapOrder
func EnumeratingMapOrder() {

	mostPopular := map[string]float32{
		"JavaScript": 65.36,
		"SQL":        49.43,
		"Python":     48.07,
		"Go":         11.15,
	}

	keys := make([]string, 0, len(mostPopular))

	for key := range mostPopular {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", mostPopular[key])
	}
}
