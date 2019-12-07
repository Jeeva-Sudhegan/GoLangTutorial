package main

import "fmt"

/*
achived using 'interface'
*/

type name interface {
	getName() string
}

type student struct {
	studentName string
}

type employee struct {
	employeeName string
}

func (s student) getName() string {
	return s.studentName
}

func (e employee) getName() string {
	return e.employeeName
}

// same function with different object parameter
// function works if any struct is called that implements name interface
// else throw error
func printName(object name) {
	fmt.Printf("The name of the %T is %v\n", object, object.getName())
}

func main() {
	s := student{"Jeeva"}
	e := employee{"Sudhegan"}
	printName(s)
	printName(e)
}
