package main

import "fmt"

type AgeT interface {
	int8 | int16
}

type NameE interface {
	string
}

type User[T AgeT, E NameE] struct {
	age  T
	name E
}

func (u *User[T, E]) GetAge() T {
	return u.age
}

func (u *User[T, E]) GetName() E {
	return u.name
}

func main() {
	var u User[int8, string]
	u.age = 18
	u.name = "weiwei"

	age := u.GetAge()
	name := u.GetName()

	fmt.Println(age, name)
}
