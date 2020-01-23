package main

import "time"

import "fmt"

func fale(pessoa, texto string, qtde int) {
	for index := 0; index < qtde; index++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração%d)\n", pessoa, texto, index+1)
	}
}

func main() {
	
}
