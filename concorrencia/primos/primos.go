package main

import "time"

import "fmt"

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primo(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 200)
				break
			}
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 30)
	go primo(cap(c), c)
	for primo := range c {
		fmt.Println(primo)
	}
	fmt.Println("Fim..")
}
