package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/tiny/emit"
	"github.com/ignite-laboratories/core/tiny/enum/direction"
	"github.com/ignite-laboratories/core/tiny/enum/traveling"
	"math/rand/v2"
)

func main() {
	var m std.Measurement
	var err error

	m, _ = emit.From[byte](77, 22, 44, 88).All()
	p := m.AsPhrase()
	fmt.Printf("\n#0 - Taking a direct binary measurement of several bytes into a phrase named %v:\n", p.Name)
	fmt.Printf("%v ← %v(byte{77}, byte{22}, byte{44}, byte{88})\n\n", p.StringPretty(), p.Name)

	random := rand.Int64()
	width := uint(17)
	m, _ = emit.From[int64](random).All()
	p = m.AsPhrase(int(width))
	fmt.Printf("#1 - Measuring a random 64 bit number into a phrase aligned at %d bits-per-measurement named %v:\n", width, p.Name)
	fmt.Printf("%v ← %v(%v)\n\n", p.StringPretty(), p.Name, random)

	fmt.Printf("#2 - Emitting from the end of %v until a condition has been met:\n", p.Name)
	width = 11 // This is a 'closure'
	continueFn := func(i uint, data []std.Bit) bool {
		if len(data) < int(width) {
			return true
		}
		return false
	}

	m, err = emit.From(p).Until(continueFn, traveling.Westbound)
	fmt.Printf("%v ← %v %v while ( len(found) < %d )\n", m.AsPhrase(-1).StringPretty(), p.Name, traveling.Westbound.StringFull(true), width)

	m, err = emit.From(p).Until(continueFn, traveling.Eastbound)
	fmt.Printf("%v ← %v %v while ( len(found) < %d )\n\n", m.AsPhrase(-1).StringPretty(), p.Name, traveling.Eastbound.StringFull(true), width)

	fmt.Printf("#3 - Emitting specific bits of %v:\n", p.Name)
	m, _ = emit.From(p).Between(11, 44)
	fmt.Printf("%v ← %v[11:44]\n\n", m.AsPhrase(-1).StringPretty(), p.Name)

	fmt.Printf("#4 - Gracefully emitting beyond the bounds of %v:\n", p.Name)
	m, err = emit.From(p).Between(55, 88)
	fmt.Printf("%v ← %v[55:88] - Error: %v\n\n", m.AsPhrase(-1).StringPretty(), p.Name, err)

	fmt.Printf("#5 -  Emitting the NOT of the last emitted bits from %v:\n", p.Name)
	notBits, _ := emit.From(m).NOT()
	fmt.Printf("%v ← !%v\n\n", notBits.AsPhrase(-1).StringPretty(), p.Name)

	fmt.Println("#6 - Measuring an object in memory:")

	m, _ = emit.From[direction.Direction](direction.Future).All()
	fmt.Printf("%v ← Measurement of [%v]\n\n", m.StringPretty(), direction.Future)

	fmt.Println("#7 - Recreating the original object from the measurement:")
	fmt.Printf("%v ← Reconstructed Object\n\n", emit.To[direction.Direction](m))

	fmt.Println("#8 - Pattern measurement:")

	pattern := []std.Bit{1, 0, 0, 1, 1}
	width = 22
	fmt.Printf("%v ← %v %v\n", emit.Pattern(width, traveling.Westbound, pattern...).AsPhrase(-1).StringPretty(), pattern, traveling.Westbound.StringFull())
	fmt.Printf("%v ← %v %v\n", emit.Pattern(width, traveling.Eastbound, pattern...).AsPhrase(-1).StringPretty(), pattern, traveling.Eastbound.StringFull())
	fmt.Printf("%v ← %v %v\n", emit.Pattern(width, traveling.Inbound, pattern...).AsPhrase(-1).StringPretty(), pattern, traveling.Inbound.StringFull())
	fmt.Printf("%v ← %v %v\n", emit.Pattern(width, traveling.Outbound, pattern...).AsPhrase(-1).StringPretty(), pattern, traveling.Outbound.StringFull())
	width = 11
	fmt.Printf("%v ← %d repeating `0`s\n", emit.Pattern(width, traveling.Westbound, 0).AsPhrase(-1).StringPretty(), width)
	fmt.Printf("%v ← %d repeating `1`s\n\n", emit.Pattern(width, traveling.Westbound, 1).AsPhrase(-1).StringPretty(), width)
}
