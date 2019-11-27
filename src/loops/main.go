package main

import (
	"fmt"
)

/*
	for is the only loop
	syntax
	for initialization; condition; post{}
	all three are optional
	semicolons can also be omitted
*/
func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // same as other lang
		}
		if i == 2 {
			continue // same as other lang
		}
		fmt.Println(i)
	}

	// label is a method used to break outer loop from the inner loop
	// e.g. is below
outer: // creating label
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			fmt.Println(i, j)
			if i == j {
				break outer // using label to break the loop
			}
		}
	}
	num := 5
	for num < 10 {
		fmt.Println(num)
		num += 2
	}
	// infinite loop
	// for {

	// }
}
