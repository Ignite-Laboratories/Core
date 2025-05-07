package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/std"
	"reflect"
)

func PrintStructure(v interface{}) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	fmt.Printf("Structure of %s:\n", t.Name())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("  %s: %s\n", field.Name, field.Type)
	}
}

func main() {
	PrintStructure(std.Data[any]{})
}
