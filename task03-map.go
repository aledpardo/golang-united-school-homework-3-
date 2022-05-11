package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	var keys []int
	for i := range input {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	var output []string
	for i := range keys {
		output = append(output, input[keys[i]])
	}
	return output
}
