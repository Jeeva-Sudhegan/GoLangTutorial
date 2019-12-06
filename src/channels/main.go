package main

import (
	"fmt"
)

/*
channels are thought as pipes through which goroutines communicate
data can be sent from one goroutine from another using channel
each channel has type that specifies that data type only should be passed through the channel
zero value of channel is nil
by default channels are pointer
chan T is channel of type T
to read from channel => data :<- channel
to write to channel => channel <- data
when sending data through channel, control is blocked until the channel data is read
vice versa also happens
This makes goroutines communication effeciently without the use of locks or conditional variables
*/

func hello(done chan bool) {
	fmt.Println("Hello from Goroutine")
	done <- true // sending data to done channel
}

func uniDirectional(data chan<- int) {
	data <- 10
}

func producer(chnl chan int) {
	for i := 0; i < 5; i++ {
		chnl <- i
	}
	close(chnl) // to close the channel
}

func main() {
	var a chan int // declaring a channel variable
	if a == nil {
		fmt.Println("Channel a is nil, defining it...")
		a = make(chan int) // defining a channel variable
		// can also use a := make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}
	done := make(chan bool)
	go hello(done)
	b := <-done // receiving data from channel. Until this happens, it won't go to next line
	// can also use like this, <- done without storing in b which is also legal
	fmt.Println("Returned from hello Goroutine", b) // we can also put <- done
	// data := make(chan<- int) // declare unidirectional flow of channel
	// go uniDirectional(data)
	// fmt.Println("data from uniDirectional function", <-data) // this won't work because we declared send only channel
	// to work in above scenario, create in bidirectional flow
	data2 := make(chan int)     // declaring bidirectional channel
	go uniDirectional(data2)    // passing it to function with unidirectional channel
	receivedData, ok := <-data2 // returns the data and boolean either the data is received or not
	if ok {
		fmt.Println("data from uniDirectional function", receivedData) // works because defined as bidirectional channel
	}
	chnl := make(chan int)
	go producer(chnl)
	for {
		receive, ok := <-chnl
		if ok == false {
			fmt.Println("Producer channel is closed")
			break
		}
		fmt.Println("From producer, data =>", receive)
	}
	// using for range loop
	chnl = make(chan int)
	go producer(chnl)
	for v := range chnl {
		fmt.Println("Inside for range data =>", v)
	}
	fmt.Println("Main function exits")
}
