package main

import "fmt"

func main() {
	funcEsalarois := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcEsalarois["Rafael Filho"] = 1350.0

	for nome, salario := range funcEsalarois {
		fmt.Println(nome, salario)
	}
}
