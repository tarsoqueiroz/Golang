package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string, 2)

	go func() {
		time.Sleep(time.Second * 3)
		messages <- "ping"
		time.Sleep(time.Second * 3)
		messages <- "ping"
	}()

	fmt.Println("wait for message 1")
	msg := <-messages
	fmt.Println("The message 1:", msg)
	fmt.Println("wait for message 2")
	msg = <-messages
	fmt.Println("The message 2:", msg)
	fmt.Println("done")

}
