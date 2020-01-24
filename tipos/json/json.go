package main

import (
	"encoding/json"
	"fmt"
)


type produto struct {
	ID    int      `json:"id"`
	Nome string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tag  []string `json:"tags"`
}

func main() {
	// struct para json
p1 := produto{1, "Notebook", 18999.90, []string{"Promocao", "Eltronico"}}

	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))

	var p2 produto


	jsonString := `{"id": 2, "nome": "Caneta", "preco": 8.9, "tags": ["Papelaria", "Importando"]}`

	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2.Tag[1])

}
