package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println(produto2.precoComDesconto())

	produto3 := produto{
		nome:     "Caneta",
		desconto: 0.05,
		preco:    1.79,
	}
	fmt.Println(produto3, produto3.precoComDesconto())

	produto4 := produto{"Desktop", 0.50, 2000.00}
	fmt.Println(produto4, produto4.precoComDesconto())
}
