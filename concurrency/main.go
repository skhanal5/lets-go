package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println(phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true // send true to the channel
	close(doneChan) // can close a channel, use when you know which operation is the slowest
}

func main() {
	done := make(chan bool) // contains a communication channel + a value to pass in the channel
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How... are... you...?", done)
	go greet("I hope you're liking the course!", done)
	for doneChan := range done {
		fmt.Println(doneChan)
	}
}
