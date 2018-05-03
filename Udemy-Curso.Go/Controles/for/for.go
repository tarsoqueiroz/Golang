package main

import "fmt"

func main() {
	fmt.Println("*** Init fora, incremento dentro")
	i := 1
	for i <= 10 { // não tem while em Go
		fmt.Printf("%d ", i)
		i++
	}

	fmt.Println("\n*** Init dentro, incremento no for")
	for i := 0; i <= 20; i += 2 { // não tem while em Go
		fmt.Printf("%d ", i)
	}

	fmt.Println("\n*** Misturando")
	for i := 0; i <= 30; i += 2 {
		fmt.Printf("%d ", i)
		if i%2 == 0 {
			fmt.Printf("é par ")
		} else {
			fmt.Printf("é impar ")
		}
		i++
	}
}
