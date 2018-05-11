package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	fmt.Println("Executar 1!")
	ch <- 1
	fmt.Println("Executar 2!")
	ch <- 2
	fmt.Println("Executar 3!")
	ch <- 3
	fmt.Println("Executar 4!")
	ch <- 4
	fmt.Println("Executar 5!")
	ch <- 5
	fmt.Println("Executar 6!")
	ch <- 6
	fmt.Println("Executou!")
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
