package main

import "time"

import "fmt"

// Channel (Canal) -  é a forma de comunicação entre goroutines
// channel é um tipo.

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando os dados

	time.Sleep((time.Second))
	c <- 3 * base

	time.Sleep((time.Second))
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("Prmeira")

	a, b := <-c, <-c // recebendo os dados do canal
	fmt.Println("Segunda e Terceira")
	fmt.Println(a, b)
	fmt.Println(<-c)
}
