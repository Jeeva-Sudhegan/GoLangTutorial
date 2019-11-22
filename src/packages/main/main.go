package main

import (
	"fmt"
	// can also use the package name as file name and import as "../package1"
	// and accessed by package1
	package12 "../package1" // <package_name> <location_of_the_package>
)

func main() {
	fmt.Println("Main function from main package")
	var area = package12.Area(10, 10)
	fmt.Println("Area is", area)
}
