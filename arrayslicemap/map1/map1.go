package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas deve ser inicializados

	aprovados := make(map[int]string)

	aprovados[3210321031] = "Maria"
	aprovados[3210321032] = "Pedro"
	aprovados[3210321033] = "Ana"

	// fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d) \n", nome, cpf)
	}

	fmt.Println(aprovados[3210321031])
	delete(aprovados, 3210321031)
	fmt.Println(aprovados)
}
