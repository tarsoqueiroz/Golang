package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 12345.67,
			"Guga Pereira":   12221.12,
		},
		"J": {
			"José João": 22111.12,
		},
		"P": {
			"Pedro Jr": 31123.31,
		},
	}
	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "P")
	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
