package main

import "time"

import "fmt"

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante
	fmt.Println("Só depois que foi lido!")
}

func main() {
	c := make(chan int) // canal sem buffer
	go rotina(c)
	fmt.Println(<-c)
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock
	fmt.Println("FIm")
}
