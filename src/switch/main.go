package main

import (
	"fmt"
)

/*
	same as other lang
*/

func main() {
	num := 5
	switch num { // switch num:=1 ; num {
	case 1:
		fmt.Println("One")
		// break 	// optional because break is automatically added at the end of case
	case 2:
		fmt.Println("Two")
		break
	case 3:
		fmt.Println("Three")
		break
	case 4:
		fmt.Println("Four")
		break
	case 6, 7, 8, 9: // multiple cases that have same output
		fmt.Println("Greater than 5")
		break
	default:
		fmt.Println("Default")
		break
	}
	switch {
	case num > 0 && num <= 5:
		fmt.Println("Inbetween 0 and 5") // it matches only this
		fallthrough                      // to continue with upcoming cases
	case num >= 5 && num < 10:
		fmt.Println("Inbetween 5 and 10")
	case num == 5: // not matched because previous case doesn't have fallthrough keyword
		fmt.Println("5")
	}
}
