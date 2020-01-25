package main

import (
	"fmt"
	"github.com/higordiego/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.youtube.com"),
		html.Titulo("https://www.google.com", "https://unileao.edu.br"),
		// html.Titulo("https://www.code3r.com.br", "https://www.google.com"),
		// html.Titulo("https://unileao.edu.br", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
