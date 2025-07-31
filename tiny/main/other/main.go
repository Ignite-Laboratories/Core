package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/tiny/enum/traveling"
	"github.com/ignite-laboratories/core/tiny/selection"
)

func main() {
	// Collect 1024 elements
	data := selection.FromWhile(func(i uint) (int, bool) {
		if i < 1024 {
			return int(i), true
		}
		return int(i), false
	})

	even := func(element int) bool {
		return element%2 == 0
	}
	under100 := func(element int) bool {
		return element < 100
	}
	above25 := func(element int) bool {
		return element > 25
	}

	// Range over the data
	for _, item := range data.Where(even).Where(under100).Where(above25).Express(traveling.Westbound).Yield {
		// Print the results
		fmt.Println(item)
	}
}
