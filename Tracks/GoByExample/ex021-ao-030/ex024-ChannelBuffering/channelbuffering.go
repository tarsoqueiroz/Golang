package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "tic"
	messages <- "toc"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
