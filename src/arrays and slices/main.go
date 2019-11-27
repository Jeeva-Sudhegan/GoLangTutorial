package main

import "fmt"

/*
	array collection of same type
	array are value types not reference type
	if a is array and b = a
	changes in b are not reflected in a
	array are not resizable
*/
func main() {
	var a [3]int // array type is also based on num of elem
	// for e.g. above a is of type [3]int
	fmt.Println(a) // initialized to 0
	a[0] = 1       // assigned and indexed same as other lang
	a[1] = 2       // not everything element should be assigned
	a[2] = 3
	fmt.Println(a)
	var b = [3]int{10, 20, 30} // short hand declaration
	fmt.Println(b)
	var c = [...]int{100, 200, 300, 400} // size of the array is omitted only when using shorthand declaration
	fmt.Println(c)
	var d = a      // copying values of a to d
	d[0] = 1000    // changing the first index value of d
	fmt.Println(a) // printing a
	a = b          // only when both variable holds same size and type
	fmt.Println(a)
	fmt.Println("Length of a", len(a))
	// range is used to iterate array
	for index, value := range a { // returns index and value of the array
		// _ is used if you want to omit any of the above values like _, value or index, _
		fmt.Println("Index", index, "Value", value)
	}
	// multidimensional array
	var x = [2][2]int{
		{1, 2},
		{3, 4}, // comma is necessary
	}
	fmt.Println("Multidimensional array", x)
	// slice same as python slice
	// powerful than array
	// flexible
	// references to existing arrays
	var z = []int{3, 4, 5, 6, 7, 8, 9}
	fmt.Println(z)
	y := z[1:4] // same syntax as python slice
	fmt.Println("Slice:", y)
	y[0] = 20
	fmt.Println(z) // change in slice reflects in the array
	// length of slice is num of elements present in slice
	// capacity is the num of elements in underlying array
	fmt.Println("Length of slice:", len(y), "Capacity of slice:", cap(y))
	arr := make([]int, 3, 3) // make([num of elem]type, len, cap) to create array and returns slice reference
	fmt.Println(arr)
	// arrays are fixed and slices are dynamic
	// slice can be used to append new elements
	// under the hood, when we append new elements to the slice,
	// new array is created and its reference is returned
	// capacity of the slice also gets doubled
	arr = append(arr, 10, 20) // to append; 10, 20 are list of args
	fmt.Println("After appending:", arr)
	// when declared array or slice without giving values and mentioning size
	// it is considered as 'nil'
	var m []int
	if m == nil {
		fmt.Println("m is nil")
		// appending is possible in m
		m = append(m, 10)
		fmt.Println("After appending m:", m)
	}
	// to join multiple slices
	var i = []int{123, 456}        // considered as slice since we left the count
	var j = []int{321, 2345, 3241} // if provided count, then it is array
	n := append(i, j...)           // should pass only slices
	fmt.Println("Created n:", n)
	// when slice is passed to function, it is pass by reference
	// so the changes are reflected
}
