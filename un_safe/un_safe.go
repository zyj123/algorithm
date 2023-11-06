package un_safe

import (
	"fmt"
	"unsafe"
)

func myUnsafe() {
	type people struct {
		weight int32
		height int
		gender int32
	}
	type people2 []int

	var xiaoming = people{
		weight: 100,
		height: 120,
		gender: 3,
	}
	//var xiaohong = people2{
	//	weight: 120,
	//	height: 120,
	//	gender: "男",
	//}

	fmt.Printf("%+v\n", *(*people2)(unsafe.Pointer(&xiaoming)))

}
