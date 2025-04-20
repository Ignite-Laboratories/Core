package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/std"
)

func main() {
	for i := 0; i < 10; i++ {
		c := std.RandomRGBA[byte]()
		fmt.Println(c)
	}
}
