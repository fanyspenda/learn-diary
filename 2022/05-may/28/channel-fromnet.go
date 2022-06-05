package main

import (
	"fmt"
)

func main() {
	// runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	go sayHelloTo(messages, "john wick")
	go sayHelloTo(messages, "ethan hunt")
	go sayHelloTo(messages, "jason bourne")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
}

func sayHelloTo(channel chan string, who string) {
	var data = fmt.Sprintf("hello %s", who)
	channel <- data
}
