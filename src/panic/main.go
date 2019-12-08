package main

import (
	"fmt"
	"runtime/debug"
)

/*
panic recover same as try catch
*/
// catch function
func recoverFunc() {
	if r := recover(); r != nil { // recover function to get the panic
		fmt.Println("Recovered from:", r)
		debug.PrintStack() // to print stack trace
	}
}

func function(a, b *string) {
	defer recoverFunc()
	defer fmt.Println("Defer in function func")
	if a == nil {
		panic("a is nil")
	} else if b == nil {
		panic("b is nil")
	} else {
		fmt.Println(a, b)
	}
}

func main() {
	defer fmt.Println("Defer from main func")
	a := "Jeeva"
	function(&a, nil)
}
