package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}
	p3 := Ponto{5.0, 5.0}

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
	fmt.Println(catetos(p1, p3))
	fmt.Println(Distancia(p1, p3))
}
