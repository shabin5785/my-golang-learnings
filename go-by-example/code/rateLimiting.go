package gobyex

import (
	"fmt"
	"time"
)

func RateLimiting() {

	//use channel and ticker to limit requests

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.NewTicker(time.Millisecond * 200)

	for req := range requests {
		<-limiter.C //block for 200 ms
		fmt.Println("request ", req, time.Now())
	}

	time.Sleep(time.Second)
	fmt.Println("burst")

	//allow short burst of requests
	//make a buffered channel of length equal the burst limit needed

	burstyLimiter := make(chan time.Time, 3)

	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}

	//every 200ms as above add to bursty limiter for burst
	go func() {
		for t := range time.NewTicker(time.Millisecond * 1000).C {
			for i := 1; i <= 3; i++ {
				burstyLimiter <- t
			}

		}
	}()

	burstyRequest := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)

	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("Request ", req, time.Now())
	}
}
