package main

import "fmt"

/*
pointer is variable that stores address of another variable
*/

func changeValByPassingPointer(point *int) {
	*point = 2
}

func changeValByPassingAddress(point *int) {
	*point = 3
}

func returnPointer() *int {
	d := 10
	return &d // allocates in heap since it going to close the local scope
}

// below method should be passed with only [3]int data type array
func arrayFunc(arr *[3]int) {
	/*
		does some operations
	*/
}

// hence we need to pass slice for functions

func main() {
	b := 255
	var a *int = &b // * indicates the accessing value from the stored address, & gets address of the variable
	// now a points to b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Printf("Address of b is %v\n", a)
	// nil will be assigned to undefined pointer variable
	var c *int
	fmt.Println("Value of c pointer is :", c)
	// creating pointer variable using 'new' keyword
	size := new(int) // new returns pointer. value of size will be the default value of the declared type here is 0
	// to be used only when using unnamed basic type
	*size = b
	fmt.Printf("Value:%d, address:%v, type:%T\n", *size, size, size)
	*size-- // won't update since new is used only for pointer reference
	fmt.Println("Value after decrementing *size:", b)
	// below will work because the created a *int type pointer variable
	*a--
	fmt.Println("Value after decrementing *a:", b)
	// passing pointer as parameter
	changeValByPassingPointer(a)
	fmt.Println("Value after calling changeValueByPassingPointer b:", b)
	// passing address as parameter
	changeValByPassingAddress(&b)
	fmt.Println("Value after calling changeValueByPassingAddress b:", b)
	// we can any of the above methods
	// return pointer from function
	getPointerFromFunc := returnPointer()
	fmt.Println("Value of the pointer returned from the function:", *getPointerFromFunc)
	// should pass slice for function call instead of arrays because of distinct
	// Go does not support pointer arithmetic that can be done in c or c++
	// a = a + 1 not possible
}
