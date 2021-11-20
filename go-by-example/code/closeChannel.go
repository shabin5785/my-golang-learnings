package gobyex

import "fmt"

func CloseChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			//second value shows if channel is closed or not
			j, more := <-jobs
			if more {
				fmt.Println("job:", j)
			} else {
				fmt.Println("recieved all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job")
	}
	close(jobs)
	fmt.Println("sent all jobs")

	//use this to wait till all jobs are done. meaning when we put mesage to done channel.
	//so block it using a wait for message read
	<-done
}
