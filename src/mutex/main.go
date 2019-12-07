package main

import (
	"fmt"
	"sync"
)

/*
	when program runs concurrently, parts of code that access shared resources should not be accessed by multiple goroutines at the same time
	this part of code is called critical section
	the output of the program depends on the sequence of execution of goroutines is called race condition
	to avoid race condition, we use mutex in go
	Mutex => locking mechanism to ensure that only one goroutine is running the critical section of code to prevent race condition
	available in 'sync' package
	two methods are there => 'Lock', 'Unlock'
	code in between 'Lock' and 'Unlock' is executed by one goroutine and not shared by others
*/
var x int = 0

func increment(wg *sync.WaitGroup, mut *sync.Mutex) {
	mut.Lock()
	x = x + 1 // other goroutines are blocked until the locked goroutine unclocks the resource
	mut.Unlock()
	wg.Done()
}

func incrementChan(wg *sync.WaitGroup, chnl chan bool) {
	chnl <- true
	x++
	<-chnl
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mut sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, &mut) // should be passed as address and not value,
		// if we passed as value, each goroutine has its own copy of mutex thus race condition still occurs
	}
	wg.Wait()
	fmt.Println("Value after incrementing using mutex:", x)
	x = 0
	chnl := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementChan(&wg, chnl)
	}
	wg.Wait()
	fmt.Println("Value after incrementing using channel:", x)
}
