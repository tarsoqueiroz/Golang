package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	fmt.Printf("Maria 1 falando...")
	fale("Maria 1", "Pq vc não fala comigo?", 3)
	fmt.Printf("João  1 falando...")
	fale("João  1", "Só posso falar depois de vc!", 1)

	fmt.Printf("Maria 2 falando...")
	go fale("Maria 2", "Ei...", 500)
	fmt.Printf("João  2 falando...")
	go fale("João  2", "Opa...", 500)

	fmt.Printf("Maria 3 falando...")
	go fale("Maria 3", "Entendi!!!", 10)
	fmt.Printf("João  3 falando...")
	fale("João  3", "Parabéns!", 5)
}
