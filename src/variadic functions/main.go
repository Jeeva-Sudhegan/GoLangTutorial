package main

import "fmt"

/*
	Function that has variable num of args
	only the last parameter should end with ...
	because variadic parameter accepts 0
	so that if we try to put the variadic parameter at first
	whatever we send will be acquired by variadic parameter without
	leaving them to other parameters.
	converts the variadic parameter to slice of that parameter type
*/
func main() {
	hello(1, 2, 3, 4, 5, 6)
	hello(1) // we can also pass 0 args
	if findNum(1, 5, 2, 3, 1, 4) {
		fmt.Println("It is present in the array")
	}
	if findNum(1, []int{1, 2, 3, 4, 5, 1}...) { // see below for explanation
		fmt.Println("It is present in the array")
	}
}

func hello(a int, b ...int) {
	fmt.Println("a:", a, "b:", b) // b will be printed as array
}

func findNum(num int, nums ...int) bool {
	for _, v := range nums {
		if v == num {
			return true
		}
	}
	return false
}

/*
	difference between slice and variadic parameter
	1. we need to create empty slice for checking num is present
	2. no need of creating slice for each call
	3. more readable
	append is a variadic function
	if we try to send slice to variable parameter then it throws compiler error
	because variable parameter expects num of args in which in turn converted to slice
	if we send slice, it will again convert to slice which is impossible
	findNum(1, []int{1,2,3,4,5,1}) is impossible
	to send slice, we can do like below
	findNum(1, []int{1,2,3,4,5,1}...)
*/
