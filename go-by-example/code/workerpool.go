package gobyex

import (
	"fmt"
	"time"
)

func WorkerPool() {

	const numOfJobs = 5
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	for w := 1; w <= 3; w++ {
		go workerFn(w, jobs, results)
	}

	for j := 1; j <= numOfJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for r := 1; r <= numOfJobs; r++ {
		fmt.Println(<-results)
	}

}

func workerFn(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("\n worker %d started job %d", id, j)
		time.Sleep(time.Second * 1)
		fmt.Printf("\n worker %d finished ", id)
		results <- j * 2
	}
}
