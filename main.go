package main

import "fmt"

func main() {
	for _, v := range nums() {
		fmt.Println(v)
	}
}

func nums() []int {
	fmt.Println("hahah###")
	return []int{1, 2, 3, 4, 5}
}
