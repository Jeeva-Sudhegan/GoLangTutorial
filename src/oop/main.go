package main

import "oop/Person"

/*
go not provide class but structs
methods can be added to structs
*/

func main() {
	person1 := Person.New("Jeeva Sudhegan", 21)
	person1.String()
}
