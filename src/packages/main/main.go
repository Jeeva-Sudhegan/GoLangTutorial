package main

import (
	"fmt"

	package12 "../package1" // <package_name> <location_of_the_package>
)

func main() {
	fmt.Println("Main function from main package")
	var area = package12.Area(10, 10)
	fmt.Println("Area is", area)
}
