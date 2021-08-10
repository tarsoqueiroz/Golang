package main

import (
	"fmt"
	"runtime/debug"
)

func f3() {
	fmt.Println("entrei no f3()")
	debug.PrintStack()
	fmt.Println("sai do f3()")
}

func f2() {
	fmt.Println("entrei no f2()")
	f3()
	fmt.Println("sai do f2()")
}

func f1() {
	fmt.Println("entrei no f1()")
	f2()
	fmt.Println("sai do f1()")
}

func main() {
	fmt.Println("entrei no main()")
	f1()
	fmt.Println("sai do main()")
}
