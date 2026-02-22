package main

import (
	"fmt"
	"slices"
)

func main() {
	m := map[string]int{"banana": 3, "apple": 5, "cherry": 2}

	//collect keys
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// sort keys
	slices.Sort(keys)

	//Iterate in order
	//The _ (blank identifier) ignores the index of the slice (0, 1, 2...).
	for _, k := range keys {
		fmt.Printf("%s:%d\n", k, m[k])
	}

}
