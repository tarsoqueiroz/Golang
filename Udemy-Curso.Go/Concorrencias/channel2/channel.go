package main

// Channel (canal) - é a forma de comunicação entre goroutines
// channel é um tipo

import (
	"fmt"
	"time"
)

func doisTresQuadroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base
}
func main() {
	ch := make(chan int)

	go doisTresQuadroVezes(2, ch)

	a, b := <-ch, <-ch // recebendo dados do canal
	fmt.Println(a, b)

	fmt.Println(<-ch)
}
