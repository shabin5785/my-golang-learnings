package gobyex

import (
	"fmt"
	"time"
)

func Channels() {

	//no buffered channel. If written outside a go routine, will
	//block main thread leading to deadlock
	messages := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		messages <- "ping"
	}()

	go func() {
		msg := <-messages
		fmt.Println("Received:", msg)
	}()

	time.Sleep(time.Second * 3)

	//buffered channel
	b := make(chan string, 2)

	b <- "one"
	b <- "two"

	fmt.Println(<-b)
	fmt.Println(<-b)
}

func ChannelsSync() {

	done := make(chan bool)
	go worker(done)
	<-done
	fmt.Println("exit")

	ping := make(chan string, 1)
	pong := make(chan string, 1)

	pings(ping, "hello")
	pongs(ping, pong)

}

func worker(done chan bool) {
	fmt.Println("Working")
	time.Sleep(time.Second * 2)
	fmt.Println("done")
	done <- true
}

//restrict channel arg to either send or receive
//can only send to pings channel
func pings(pings chan<- string, msg string) {
	pings <- msg
	// <-pings error
}

func pongs(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func SelectIt() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func RangeChannel() {
	q := make(chan string, 2)
	q <- "one"
	q <- "two"
	close(q)
	//can close an non empty channel
	// can receive from a closed channel

	for e := range q {
		fmt.Println(e)
	}
}
