package main

import "fmt"

/*
defer function is executed when the calling function is completed or returned
below example function() is executed after printing the largest number
it is also for methods
parameters in deferred function is not changed though the function is executed last
if function has multiple defer functions, executed in LIFO order
*/

func function() {
	fmt.Println("called Function")
}

type person struct {
	name string
}

func (p person) getName() {
	fmt.Println("Name is", p.name)
}

func print(str string) {
	fmt.Println("Value in deferred function", str)
}

func largest(nums []int) {
	defer function()
	p := person{"Jeeva"}
	defer p.getName() // executed in LIFO
	largest := 0
	for _, v := range nums {
		if largest < v {
			largest = v
		}
	}
	fmt.Println("Largest number is", largest)
}

func main() {
	str := "Sudhegan"
	defer print(str)
	largest([]int{3, 5, 4, 2, 1, 9, 8, 6})
	str = "Naruto"
	fmt.Println("Value before deferred function", str)
}
