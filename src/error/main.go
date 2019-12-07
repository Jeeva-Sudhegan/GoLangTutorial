package main

/*
the last return value of any function or method is error
if error is nil, there is no error
else there is an error
*/

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened!")
}
