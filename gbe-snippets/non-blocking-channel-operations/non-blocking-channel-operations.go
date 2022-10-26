package main

import "fmt"

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Message received: ", msg)
	default:
		fmt.Println("No message received")
	}

	msg := "here!"
	select {
	case messages <- msg:
		fmt.Println("Message received: ", msg)
	default:
		fmt.Println("No message received")
	}

	select {
	case msg := <- messages:
		fmt.Println("Message received ", msg)
	case sig := <- signals:
		fmt.Println("Signal received ", sig)
	default:
		fmt.Println("Nothing received")
	}

}