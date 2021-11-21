package gobyex

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(i)
		i := i //closure variable capture fix

		go func() {
			defer wg.Done()
			workerWG(i)
		}()

	}

	wg.Wait()
}

func workerWG(id int) {
	fmt.Printf("\n worker %d starting \n", id)
	time.Sleep(time.Second)
	fmt.Printf("\n worker %d finished \n", id)

}
