package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio(number int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(number)
}

func main() {
	if i := numeroAleatorio(10); i > 5 {
		fmt.Println("Ganou!!")
	} else {
		fmt.Println("Perdeu!!")
	}
}
