package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)
	fmt.Println(aprovados)

	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95135745682] = "Ana"
	aprovados[12345678978] = "Mariana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[95135745682])
	delete(aprovados, 95135745682)
	fmt.Println(aprovados[95135745682])
	fmt.Println(aprovados)
	fmt.Println(aprovados[95135745682])
	fmt.Println("----------------")
}
