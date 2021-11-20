package gobyex

import (
	"fmt"
	"time"
)

func Timer() {
	timer1 := time.NewTimer(time.Second * 2)

	//default channel provided by timer to notify
	c := <-timer1.C
	fmt.Println("timer 1 fired ", c)

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer2 fired")
	}()
	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("timer2 stopped")
	}

	//need to sleep as timer stop can be async
	time.Sleep(time.Second * 2)
}
