package gobyex

import "fmt"

func BetterChannel() {

	messages := make(chan string)
	signals := make(chan bool)

	//non blocking receive
	select {
	case msg := <-messages:
		fmt.Println("message ", msg)
	default:
		fmt.Println("no message")
	}

	msg := "hi"

	//non blocking send
	select {
	case messages <- msg:
		fmt.Println("Sent")
	default:
		fmt.Println("Nothing to sent")
	}

	//non blocking multi channel
	select {
	case msg := <-messages:
		fmt.Println("message ", msg)
	case msg := <-signals:
		fmt.Println("signal ", msg)
	default:
		fmt.Println("no message")
	}

}
