package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("Zeroes from:\n - int: %v\n - float64: %v\n - bool: %v\n - string: %q\n - pointer: %v", a, b, c, d, e)
}
