package main

import "oop/person"

/*
go not provide class but structs
methods can be added to structs
*/

func main() {
	person1 := person.New("Jeeva Sudhegan", 21)
	person1.String()
}
