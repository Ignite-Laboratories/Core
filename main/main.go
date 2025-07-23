package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"regexp"
)

func main() {
	avg := 0
	for i := 0; i < 1<<10; i++ {
		count := 0

		name := core.RandomNameFiltered(func(name core.GivenName) bool {
			var nonAlphaRegex = regexp.MustCompile(`^[a-zA-Z]+$`)
			return nonAlphaRegex.MatchString(name.Name)
		})

		fmt.Println(name, "|", name.Details)

		avg += count
	}
	avg /= 1 << 10
	fmt.Println(avg)
}
