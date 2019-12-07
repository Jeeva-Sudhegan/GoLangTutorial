package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// /*
// Channels are basically unbuffered
// sending data in unbuffered channels are blocking
// possible to create a channel with buffer
// sending to buffered channel is blocking only if the buffer is full
// receiving from buffered channel is blocking only if the buffer is empty
// created by adding capacity parameter in make()
// syntax:
// 	channel_name := make(chan <type>, <capacity>)

// WaitGroup is nothing but main func wait till all the goroutines are completed

// Worker pool is a collection of threads which are waiting for task to be assigned to them
// after completing the task, they make themselves available for next
// */

// func print(chnl chan int) {
// 	for i := 0; i <= 6; i++ {
// 		fmt.Println("Passed", i, "to the channel")
// 		chnl <- i // after passing 2, it is blocked and control goes to main function
// 		// after one receving is done, comes back here
// 	}
// 	close(chnl)
// }

// // WaitGroup example
// func process(i int, wg *sync.WaitGroup) {
// 	fmt.Println("Started Goroutine", i)
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Ended Goroutine", i)
// 	wg.Done() // removes the process
// }

// // worker pool example
// type job struct {
// 	id        int
// 	randomnum int
// }
// type result struct {
// 	job job
// 	sum int
// }

// // worker pool channel
// var jobs = make(chan job, 10)
// var results = make(chan result, 10)

// func digits(num int) int {
// 	sum := 0
// 	no := num
// 	for no != 0 {
// 		digit := no % 10
// 		sum += digit
// 		no /= 10
// 	}
// 	time.Sleep(2 * time.Second)
// 	return sum
// }

// // worker goroutine
// func worker(wg *sync.WaitGroup) {
// 	for job := range jobs {
// 		output := result{job, digits(job.randomnum)}
// 		results <- output
// 	}
// }

// // worker pool
// func workerPool(numOfWorker int) {
// 	var wg sync.WaitGroup
// 	for i := 0; i < numOfWorker; i++ {
// 		wg.Add(1)
// 		go worker(&wg)
// 	}
// 	wg.Wait()
// 	close(results)
// }

// // allocate job
// func allocate(numOfJobs int) {
// 	for i := 0; i < numOfJobs; i++ {
// 		randomnum := rand.Intn(999)
// 		job := job{i, randomnum}
// 		jobs <- job
// 	}
// 	close(jobs)
// }

// //result
// func output(done chan bool) {
// 	for result := range results {
// 		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomnum, result.sum)
// 	}
// 	done <- true
// }

// func main() {
// 	// chnl := make(chan int, 2)
// 	// go print(chnl)
// 	// for v := range chnl {
// 	// 	fmt.Println("Received", v, "from channel")
// 	// 	time.Sleep(2 * time.Second)
// 	// }
// 	// var wg sync.WaitGroup
// 	// for i := 0; i <= 3; i++ {
// 	// 	wg.Add(1)          // num of processes, 1 denotes how many process going to spawn at that iteration
// 	// 	go process(i, &wg) // calling 1 process so wg.Add(1)
// 	// }
// 	// wg.Wait() // wait till all the process completed
// 	// fmt.Println("End of WaitGroup")

// 	// worker pool example
// 	fmt.Println("Worker pool example")
// 	startTime := time.Now()
// 	noOfJobs := 20
// 	go allocate(noOfJobs)
// 	done := make(chan bool)
// 	go output(done)
// 	noOfWorkers := 10
// 	workerPool(noOfWorkers)
// 	<-done
// 	endTime := time.Now()
// 	diff := endTime.Sub(startTime)
// 	fmt.Println("total time taken ", diff.Seconds(), "seconds")

// }

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		fmt.Println("Passed", output, "to results")
		results <- output
	}
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		fmt.Println("Passed", job, "to jobs")
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}
func main() {
	startTime := time.Now()
	noOfJobs := 30
	fmt.Println("Started allocating...")
	go allocate(noOfJobs)
	fmt.Println("Finished allocating...")
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
