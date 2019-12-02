package main

import "fmt"

/*
defines behaviour of the object
holds what the object should do
the way of implementing is upto the object
set of method signatures
same as OOPs
*/

// syntax
// type <name> interface {}

// internally considered as (type, value)
// type => concrete type that implements this interface
// value => value of the object that implements this interface
type vowelsFinder interface {
	findVowels() []rune
}

type myStr string

func (str myStr) findVowels() []rune {
	var vowels []rune
	for _, each := range str {
		switch each {
		case 'a', 'e', 'i', 'o', 'u':
			vowels = append(vowels, each)
		}
	}
	return vowels
}

// using method defined in interface makes the receiver implements the interface implicitly
// the type is considered as implemented only if all the methods in interface are implemented

// power of interface in Go

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

//salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

/*
total expense is calculated by iterating though the SalaryCalculator slice and summing
the salaries of the individual employees
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary() // calls respective type's CalculateSalary method
	}
	fmt.Printf("Total Expense Per Month $%d\n", expense)
}

type empty interface{} // empty interface, everything will implement this interface

func definition(e empty) {
	fmt.Printf("Type is %T and value is %v\n", e, e)
}

// Type assertion is used to extract underlying value of the interface
// <interface>.(T) is the syntax

func assertInt(e empty) {
	s, ok := e.(int) // to run even if other type are passed
	if ok {          // ok is true only if the passed arg is of type int
		fmt.Println(s)
	} else {
		fmt.Println("Not int")
	}
}

// Type Switch
// same as switch case
// except the cases are type instead of values

func findType(e empty) {
	switch e.(type) {
	case vowelsFinder:
		fmt.Println("vowelsFinder interface type")
	case string:
		fmt.Println("I am string")
	case int:
		fmt.Println("I am int")
	case float64:
		fmt.Println("I am float64")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	name := myStr("Jeeva Sudhegan")
	var v vowelsFinder
	v = name                                          // type => myStr, value => "Jeeva Sudhegan"
	fmt.Printf("The vowels are %c\n", v.findVowels()) // name.findVowels() also work

	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1} // since both permanent and contract implements interface can be casted
	// even if we create new type and add to the above slice, nothing is changed in the below function
	totalExpense(employees)
	definition("Hello")
	definition(1)
	definition(2.3)
	assertInt(2) // runs only if the type is 'int', other than 'int' throws runtime error
	findType("Hello")
	findType(123)
	findType(2.34)
	findType(name) // comparing the interface with the type that implements
}
