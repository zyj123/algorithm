package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {
	fmin := GMin[int]

	x := fmin(2, 3)
	fmt.Printf("%T", x)
}

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
