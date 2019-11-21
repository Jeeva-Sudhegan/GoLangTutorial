package main

import "fmt"

func main() {
	fmt.Println(function("Jeeva", "Sudhegan", 105))
	var area, perimeter float64
	area, perimeter = rectPropsWithoutNamedReturnVariables(10, 10)
	fmt.Println("From WithoutNamedReturnVariables area", area, "perimeter", perimeter)
	area, perimeter = rectPropsWithNamedReturnVariables(10, 10)
	fmt.Println("From WithNamedReturnVariables area", area, "perimeter", perimeter)

	_, perimeter = rectPropsWithoutNamedReturnVariables(20, 20)      // use _ if you need to discard that value from returned values
	fmt.Println("received only perimeter ", perimeter, "area", area) // prev area value is displayed

}

func function(param1 string, param2 string, param3 int) string { // or func function(param1 , param2 string, param3 int)
	return "Hello from " + param1 + " " + param2 + string(param3) // should type cast explicitly
}

func rectPropsWithoutNamedReturnVariables(length, breadth float64) (float64, float64) { // multiple return variables
	var area = length * breadth
	var perimeter = 2 * (length + breadth)
	return area, perimeter
}

func rectPropsWithNamedReturnVariables(length, breadth float64) (area, perimeter float64) { // named return variables
	area = length * breadth
	perimeter = 2 * (length + breadth)
	return
}
