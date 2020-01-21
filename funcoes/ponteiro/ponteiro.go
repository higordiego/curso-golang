package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por um *

func inc2(n *int) {
	// revisão * é usado para desreferenciar
	// ter acesso ao valor no qual o ponteiro aponta.
	*n++
}

func main() {
	n := 1
	inc1(n)
	fmt.Println(n) // por valor - copia

	inc2(&n) // por referência
	fmt.Println(n)

}
