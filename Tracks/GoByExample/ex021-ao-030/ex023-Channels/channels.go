package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)

	go func() {
		time.Sleep(time.Second * 10)
		messages <- "ping"
		time.Sleep(time.Second * 10)
	}()

	fmt.Println("wait for message")
	msg := <-messages
	fmt.Println("The message:", msg)
	fmt.Println("done")

}
