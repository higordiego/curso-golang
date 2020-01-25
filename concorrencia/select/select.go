package main

import "github.com/higordiego/html"

import "time"

import "fmt"

func osMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := osMaisRapido(
		"https://www.cod3r.com.br",
		"https://www.youtube.com",
		"https://www.google.com",
	)
	fmt.Println(campeao)
}
