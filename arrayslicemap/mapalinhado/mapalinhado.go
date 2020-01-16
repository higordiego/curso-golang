package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela": 12321312.20,
			"Guga":     321321321.2,
		},
		"J": {
			"José": 123.50,
		},
		"P": {
			"Pedro Junior": 3213.30,
		},
	}

	delete(funcPorLetra, "P")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra)
		for name, value := range funcs {
			fmt.Println(name, value)
		}
	}
}
