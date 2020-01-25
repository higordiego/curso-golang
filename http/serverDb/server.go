package main

import "net/http"

import "log"

func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	log.Printf("executando")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
