package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// fmt.Println("Ovalor de x é" + x) ** GERA ERRO **
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)
	fmt.Printf("O valor de x e %f\n", x)
	fmt.Printf("O valor de x é %.2f\n", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("%d  %f  %.1f  %t  %s\n", a, b, b, c, d)
	fmt.Printf("%v  %v  %v  %v\n", a, b, c, d)

}
