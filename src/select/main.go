package main

import (
	"fmt"
	"time"
)

/*
select statement used to choose from multiple send/receive channel operations
blocks until one of the send/receive operation is ready
if multiple are ready, one is chosen random
syntax similar to switch, except each case will be channel operation
*/

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	go server1(chan1)
	go server2(chan2)
	select {
	case s1 := <-chan1:
		fmt.Println(s1)
	case s2 := <-chan2:
		fmt.Println(s2)
	default:
		fmt.Println("No value received")
	}
	// select {}  // results in deadlock because select will block until read/write in channel is happened
}
