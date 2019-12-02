package main

import "fmt"

/*
special type of functions
that has receiver type
we can link those functions to
respective structs.
to achieve class like functionality
same method name for different types but same function name not possible
function receives either pointer or value
method receivers receive both value and pointer
can also created for other than struct types
*/

// syntax
// func (t Type) methodName(parameter list) {}

type myNumber int // should create alias because int is in different package
// the rule of method is both the declaration and method should be in same package
// int is in different package, so we created alias for int
// NOTE: though myNumber is alias of int, int and myNumber are different to GoLang
// e.g.
// var a int = 5
// var b myNumber = a
// above two lines will throw error because a and b are different types according to Go

func (a myNumber) add(b myNumber) myNumber {
	return a + b
}

type Student struct {
	name string
	age  int
	address
}

type address struct {
	city  string
	state string
}

func (a address) fullAddress() { // has the anonymous field as receiver
	fmt.Println(a.city, a.state)
}

func (student Student) setName(name string) { // receive value receivers
	student.name = name
}

func (student *Student) setAge(age int) { // receive pointer receivers implicitly
	student.age = age
}

func main() {
	student := Student{
		"Jeeva",
		21,
		address{
			city:  "Chennai",
			state: "Tamil Nadu",
		},
	}
	var a, b myNumber = 5, 10
	student.setName("Naruto") // called by value type, so changes are not reflected
	fmt.Println("Student:", student)
	student.setAge(22) // called by pointer type, so changes are reflected
	fmt.Println("Student:", student)
	student.fullAddress() // called directly using struct that contains anonymous field
	fmt.Println("Addition using method")
	fmt.Println("a.add(b):", a.add(b))
}
