package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(reflect.TypeOf(a1))
	fmt.Println(reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	s2 := a2[1:3]

	fmt.Println(a2)
	fmt.Println(s2)

	s3 := a2[:2]

	fmt.Println(s3)

	s4 := s2[:1]
	fmt.Println(s2, s4)
}
