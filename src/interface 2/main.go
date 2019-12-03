package main

import "fmt"

/*
in interface 1 implemented value receivers
this implements pointer receivers
zero value of interface is nil
calling any interface method will nil, will occur `panic` error
*/

type describer interface {
	describe()
}

type giveName interface {
	getName() string
}

// we can compose different interface under single interface
type personOperations interface {
	describer
	giveName
}

type person struct {
	name string
}

type address struct {
	city string
}

// single type can implement multiple interfaces like below
// person type implements describer and giveName interfaces
func (p person) describe() {
	fmt.Println(p)
}

func (p person) getName() string {
	return p.name
}

func (a *address) describe() {
	fmt.Println(a)
}

func main() {
	p := person{
		"Jeeva",
	}
	var d1 describer
	d1 = p
	d1.describe()
	fmt.Println("Using giveName interface name:", p.getName())
	d1 = &p
	d1.describe()
	a := address{
		"chennai",
	}
	// d1 = a
	// d1.describe()
	// throws error because we implemented describer in address as pointer receiver not as value receiver
	// in interface, if value receiver is used to implement the interface, both value and pointer receiver is used.
	// if pointer receiver is used, only pointer receiver is used
	// it is because the concrete value stored in interface is not addressable
	d1 = &a
	d1.describe()
	var personOp personOperations = p // create variable of type interface and assign the object to it or directly you can access
	fmt.Println("Person name:", personOp.getName())
	personOp.describe()
}
