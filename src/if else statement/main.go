package main

import (
	"fmt"
)

/*
	{} are mandatory for if else statements
	no () around condition
	can also has another variant in which declared and check condition

*/

func main() {
	num := 10
	if num%2 == 0 { // if num:=10; <condition> {
		// if used like in above comment, then num cannot be accessed outside if else block
		fmt.Println("Even")
	} else { // else or else if should start in the closing parenthesis line or it gives error
		// it is because semi colon is added automatically to the end of the } parenthesis
		fmt.Println("Odd")
	}
}
