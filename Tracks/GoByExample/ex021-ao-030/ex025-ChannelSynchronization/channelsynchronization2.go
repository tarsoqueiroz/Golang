package main

import (
	"fmt"
	"time"
)

func workerTurtle(doneTurtle chan bool) {
	fmt.Println("Turtle working...")
	time.Sleep(time.Second * 5)
	fmt.Println("Turtle done!!!")

	doneTurtle <- true
}

func workerRabbit(doneRabbit chan bool) {
	fmt.Println("Rabbit working...")
	time.Sleep(time.Second)
	fmt.Println("Rabbit done!!!")

	doneRabbit <- true
}

func main() {

	doneTurtle := make(chan bool, 1)
	doneRabbit := make(chan bool, 1)

	go workerTurtle(doneTurtle)
	go workerRabbit(doneRabbit)

	fmt.Println("Are they working?")
	<-doneTurtle
	fmt.Println("From Turtle OK.")
	<-doneRabbit
	fmt.Println("From Rabbit OK.")
	fmt.Println("They worked.")

}
