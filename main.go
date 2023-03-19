package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(9))
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(0))
	fmt.Println(mySqrt(10))
	fmt.Println(mySqrt(16))
	fmt.Println(mySqrt(15))
	fmt.Println(mySqrt(17))
}
func mySqrt(x int) int {
	l, r := 0, x
	ans := 0
	for l <= r {
		m := l + (r-l)/2
		if m*m <= x {
			ans = m
			l = m + 1
		} else if m*m > x {
			r = m - 1
		}
	}
	return ans
}
