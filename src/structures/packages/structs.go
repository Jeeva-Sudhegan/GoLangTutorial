package packages

/*
user defined type
collection of fields
group data into single unit
*/

// declaring structure
// syntax
// type <struct_name> struct {
// 	<var_name> <var_type>
// 	....
// }

// to export structs, capitalize starting words
// also the fields so that they can be accessed from packages

type Student struct { // named structure because contains name
	FirstName, LastName string
	Age                 int
}

// Nested Structs
type Address struct {
	City, State string
}

type Employee1 struct {
	Name    string
	Address Address
}
