package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // não tem while em Go
		fmt.Printf("%d ", i)
		i++
	}

	fmt.Println("\ni pos 1ºfor: ", i)

	for i := 0; i <= 20; i += 2 {
		fmt.Println("i = ", i)
	}

	fmt.Println("i pos 2ºfor: ", i)

	fmt.Println("\nMisturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	fmt.Println("i pos 3ºfor: ", i)

	for {
		// laço infinito
		fmt.Println("Para sempre a cada 30s...")
		time.Sleep(time.Second * 3)
	}

	// Veremos o foreach no capítulo de array
}
