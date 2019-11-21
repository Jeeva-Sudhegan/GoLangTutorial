package main

import (
	"fmt"
	"unsafe"
)

/*
	bool
	Numeric Types
		int8, int16, int32, int64, int
		uint8, uint16, uint32, uint64, uint
		float32, float64
		complex64, complex128
		byte
		rune
	string
*/

func main() {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)
	e := "string"
	fmt.Println("String e ", e)
	f := 20
	fmt.Printf("The type is %T and size is %d\n", f, unsafe.Sizeof(f))
	g := 10.10
	fmt.Printf("The type is %T and size is %d\n", g, unsafe.Sizeof(g))
	comp := complex(1, 1)
	fmt.Println("Complex number is ", comp)

	/*
		no automatic or implicit conversion of data types
		should change explicitly
		eg
		x := 10
		y := 20.20
		sum := x + y      //not possible because using variables mismatch */
	sum := 10 + 20.20 // constants can be used
	fmt.Println(sum)
	/*sum := x + int(y) //is possible
	this is same for assignment too
	*/
}
