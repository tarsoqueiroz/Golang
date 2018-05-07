package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo pois já estão definido na função
	// return segundo, primeiro // não seria necessário explicitar assim
	// return primeiro, segundo // poderia fazer assim sobrepondo a definição

}

func main() {
	x, y := trocar(1, 2)
	fmt.Println(x, y)

	segundo, primeiro := trocar(3, 4)
	fmt.Println(segundo, primeiro)
}
