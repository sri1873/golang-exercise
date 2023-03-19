package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	i int
	f float64
	b bool
	s string
)

func main() {
	fmt.Printf("Type: %T Value: %t\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %d\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("%d %g %t %q\n", i, f, b, s)
}
