package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // zero-valued variable
	fmt.Println(e)

	f := "apple" // shorthand declare and initializing
	fmt.Println(f)
}
