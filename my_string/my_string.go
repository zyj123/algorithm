package my_string

import (
	"fmt"
	"unsafe"
)

func myString2() {
	b1 := []byte{'h', 'e', 'l', 'l', 'o'}
	s := *(*string)(unsafe.Pointer(&b1))

	m := map[string]string{}
	m[s] = "hhh"

	fmt.Println(s, m["hello"])

	b1[0] = 'w'

	fmt.Println(s, m["hello"])
}

func myString() *string {
	//b := []byte{'1', 'a', 'v', 'b', '3'}
	//
	//fmt.Println(*(*string)(unsafe.Pointer(&b)))

	b1 := []byte{'h', 'e', 'l', 'l', 'o'}
	b1[0] = 'w'

	fmt.Println(b1)

	s := "hello"
	b2 := *(*[]byte)(unsafe.Pointer(&s))
	b2[0] = 'w'
	fmt.Println(b2)

	s2 := s + "world"
	b3 := *(*[]byte)(unsafe.Pointer(&s2))
	b2[0] = 'w'
	fmt.Println(b3)

	return &s
}
