package main

import (
	"fmt"
	"time"
)

/*
functions or methods run concurrently with other functions or methods
light weight threads
cost of creating goroutines compared to threads is very less
so it is common for having 1000s of goroutines
Advantages of Goroutine over Thread:
	1. cheap compared to threads. memory can be allocated when application needs whereas threads need to specified
	2. Goroutines are multiplexed to fewer number of OS thread.
		One thread with thousands of goroutines.
		if one goroutine blocks, new thread is created and other goroutines are moved to the new one.
	3. communicate through channels. These channels prevent race conditions when accessing shared memory.
		channels are thought as pipe through which goroutines communicate
*/

func goroutine() {
	fmt.Println("Printed from Goroutine")
}

func numbers1() {
	for i := 0; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println("Number1 is", i)
	}
}

func numbers2() {
	for i := 5; i < 10; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Number2 is", i)
	}
}

func main() {
	go goroutine() // call the function with prefixing `go`
	// not prints the message if comment the line 47
	// function call with prefixed with `go` won't wait to return the control
	// immediately the next line is executed
	// if the main method terminated, then the return of goroutine is ignored.
	// to see the output of goroutine, we need to make the main in running
	go numbers1()
	go numbers2()
	time.Sleep(2 * time.Second) // sleep method to see the execution of the goroutine
	// the above is used as hack for checking how goroutine is working
	// instead of above, channels can be used to block
	fmt.Println("Message from main function")
}
