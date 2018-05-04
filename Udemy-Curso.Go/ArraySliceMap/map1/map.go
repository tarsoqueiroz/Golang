package main

import "fmt"

func main() {
	//	var aprovados map[int]string  // mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123] = "Maria"
	aprovados[456] = "Pedro"
	aprovados[789] = "João"

	fmt.Println(aprovados)
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	delete(aprovados, 456)
	fmt.Println(aprovados[456])
	fmt.Println(aprovados)

}
