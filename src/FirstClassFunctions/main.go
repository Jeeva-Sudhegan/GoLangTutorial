package main

import "fmt"

/*
same as higher order functions in javascript
assign function to variable
pass function as parameter
return function from function or method
*/

// higher order function
func simple(function func()) func() {
	function()
	return func() {
		fmt.Println("Function returned from simple func")
	}
}

func main() {
	a := func() { // also called anonymous function since they don't have name
		fmt.Println("Function with variable a")
	}
	// to call this function, use the variable name
	a()
	fmt.Printf("The type of a is %T\n", a)
	func() {
		fmt.Println("IIFD function")
	}() // IIFD without outer () in function definition
	func(n string) {
		fmt.Println("Passed string is:", n)
	}("Jeeva Sudhegan") // passing parameter in IIFD
	b := func() {
		fmt.Println("Example of Higher Order Functions")
	}
	simple(b)() // passing function as parameter and executing directly
}
