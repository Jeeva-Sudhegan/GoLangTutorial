package Person

import "fmt"

type Person struct {
	firstName string
	age       int
}

func New(firstName string, age int) (person Person) {
	person = Person{firstName, age}
	return
}

func (person *Person) String() {
	fmt.Println("First Name =>", person.firstName, ", age =>", person.age)
}
