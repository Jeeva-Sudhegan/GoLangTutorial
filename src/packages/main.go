package main

import (
	"fmt"

	"packages/package1"
)

/*
	returns error if imported package is not used
	to avoid these issues, use _ technique
	e.g. _ "<package name>" so that they can't be used
*/

func main() {
	fmt.Println("Main function from main package")
	var area = package1.Area(10, 10)
	fmt.Println("Area is", area)
}
