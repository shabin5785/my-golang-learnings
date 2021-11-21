package gobyex

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func AtomicCounter() {

	// to avoid main from deadlock
	go atomicFn()
	time.Sleep(time.Second * 2)
	//sleep to allow routine to finish
}

//atomic int and waitgroup can be used to sync go routines
//also to manage state between go routines
func atomicFn() {
	var ops uint64 //always positive int

	var wg sync.WaitGroup

	for i := 1; i <= 50; i++ {
		wg.Add(1)

		//each routine increments atomic int by 1000
		go func() {
			for c := 1; c <= 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println("atomic ", ops)
}
