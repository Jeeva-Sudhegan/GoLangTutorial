package main

import (
	"fmt"
	"structures/packages"
)

// we can create anonymous fields in the struct
type Anonymous struct {
	string
	bool
}

// anonymous nested struct
type Employee2 struct {
	Name string
	packages.Address
}

func main() {
	// creating structures
	student1 := packages.Student{
		FirstName: "Jeeva Sudhegan",
		LastName:  "G",
		Age:       21,
	}
	student2 := packages.Student{"Hello", "world", 10} // define in the order of declared fields
	fmt.Println("Student1:", student1)
	fmt.Println("Student2:", student2)

	// anonymous structure
	employee := struct { // struct variable is created not the struct type so that we cannot reuse
		name   string
		salary int
	}{
		name:   "employee",
		salary: 20000,
	}
	fmt.Println("Anonymous struct:", employee)
	// creating struct variable without initializing
	var student3 packages.Student
	fmt.Println("Uninitialized struct:", student3) // display zero value(default value) of the fields
	// also possible like initializing some fields instead of all

	// to access <struct_var>.<field>
	fmt.Println("First name of Student1:", student1.FirstName)
	// we can also update struct variable field
	// by <struct_var>.<field> = <value>
	student1.LastName = "Govindarajan"
	fmt.Println("After update Student1:", student1)

	// we can also use pointers in struct
	student4 := &packages.Student{
		FirstName: "student4",
		LastName:  "last",
		Age:       19,
	}
	fmt.Println("Student4:", *student4)
	// to access fields of the pointer struct
	fmt.Println("Student4 firstname:", (*student4).FirstName)
	// has option to access directly (implicit dereference)
	fmt.Println("Student4 lastname:", student4.LastName)

	// to use anonymous fields in struct
	anonymous := Anonymous{
		"Name",
		true,
	}
	// only able to define exact fields
	// if we try to add another string above, throws error that it is duplicate
	fmt.Println("Anonymous fields in struct:", anonymous)
	// to access anonymous fields
	fmt.Println("String field in anonymous field struct:", anonymous.string)

	// to create nested structures
	employee1 := packages.Employee1{
		"Name",
		packages.Address{
			"chennai",
			"tamil nadu",
		},
	}
	fmt.Println("Nested Structures:", employee1)

	// promoted fields
	// this means that while nesting structs
	// if we anonymously create a nested struct field
	// the fields inside that anonymous struct
	// will be promoted to the parent struct
	// and they will be accessed as if they were in parent struct
	employee2 := Employee2{
		"Sudhegan",
		packages.Address{
			"Palakad",
			"Kerala",
		},
	}
	// since the address field is anonymous
	// fields inside the address are promoted to employee2 struct
	fmt.Println("State of employee2:", employee2.State)

	// to check two structs equal
	// they should have same fields
	// value of each field should be same
	// only compares if they hold fields that can be compared like basic types
	employee3 := Employee2{
		"Sudhegan",
		packages.Address{
			"Palakad",
			"Kerala",
		},
	}
	if employee2 == employee3 {
		fmt.Println("employee2 and employee3 are same")
	} else {
		fmt.Println("employee2 and employee3 are not same")
	}
}
