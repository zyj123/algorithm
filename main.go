package main

import "fmt"

func main() {

	fmt.Println('3' - '0')
}

func decrypt(code []int, k int) []int {
	ret := make([]int, len(code))
	if k == 0 {
		return ret
	}
	
}
