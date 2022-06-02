package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string, times int) {
		for i := 0; i < times; i++ {
			fmt.Println(msg, i)
			time.Sleep(time.Second)
		}
	}("going", 3)

	time.Sleep(time.Second * 3)
	fmt.Println("done")
}
