package gobyex

import (
	"fmt"
	"time"
)

func Ticker() {
	//tickers are for repeated and timer for once
	ticker := time.NewTicker(time.Second * 1)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick ", t)
			}
		}
	}()

	//enough sleep to see ticker working
	time.Sleep(time.Second * 10)
	ticker.Stop()
	done <- true
	fmt.Println("done")

}
