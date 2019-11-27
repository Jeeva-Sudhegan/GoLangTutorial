package main

import (
	"fmt"
)

var packageVar int // package variables same as global variables

func main() {

	var variable int // declaring variable of syntax; var <name> <type>

	fmt.Println("Hello Jeeva", variable) // variable default value is displayed 0

	variable = 20 // assignment operation
	fmt.Println("Changed variable to ", variable)

	var dec int = 10 // defining variable value while declaring
	fmt.Println("Printing defined value ", dec)

	var implicitType = 20
	fmt.Println("Type coersion ", implicitType)

	/*
		multiple variable declaration
		int can also be dropped
		var a,b = 10, 20
	*/

	var a, b int = 10, 20
	fmt.Println("Multiple variable declaration ", a, b)

	/*
		for different types of variable declaration and definition at the same time
	*/

	var (
		x = 10
		y = 20.10
	)
	fmt.Println("x:", x, " y:", y)

	/*
		we can use shorthand assignment without using var and any type
		need initial value for using this format
		accepted if one of the variables used in assignment should be same
		eg. a,b := 10, 20
		b, c := 29, 30 works because c is new
		b, c := 50, 100 won't work because both b and c are declared already
	*/

	m, n := 10, "name"
	fmt.Println("M:", m, "n:", n)

}
