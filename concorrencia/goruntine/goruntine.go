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
	// fale("Maria", "Pq vc não fala comigo?", 3)
	// fale("João", "Só posso falar depois de você?", 1)

	// go fale("Maria", "Ei....", 5000)
	// go fale("João", "Opa....", 5000)

	go fale("Maria", "Entendi!!", 10)
	fale("João", "Parabéns!!", 5)
}
