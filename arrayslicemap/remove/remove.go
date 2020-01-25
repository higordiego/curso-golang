package main

import "fmt"

func main() {
	a := []string{"A", "B", "C", "D", "E"}

	fmt.Println(a[:len(a)-1])
}
