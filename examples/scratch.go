package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
)

func main() {
	core.SetRandomNumberGenerator[byte](BadGenerator)
	for i := 0; i < 10; i++ {
		c := std.RandomRGBA[byte]()
		n := std.NormalizeRGBA32(c)
		o := std.ScaleToTypeRGBA32[byte](n)
		fmt.Println(c, n, o)
	}
}

func BadGenerator[T core.Numeric](r core.Tuple[T]) T {
	return 7
}
