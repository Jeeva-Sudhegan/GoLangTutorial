package package1

import "fmt"

/*
	to export the function write the function name with caps at start
	eg. Area, Perimeter, etc.
*/

/*
	Constructor like function that is automatically executed when this package is imported in another package
	should not have any return type and parameters
	cannot called explicitly
	perform initialization tasks
	can have multiple init functions and executed in the order they have implemented
*/
func init() {
	fmt.Println("Hello from package 1")
}

func Area(length, breadth float64) (area float64) {
	area = length * breadth
	return
}

func Perimeter(length, breadth float64) (perimeter float64) {
	perimeter = 2 * (length + breadth)
	return
}
