package main

import "fmt"

func main() {
	const a = 10
	fmt.Println("Constant a:", a)
	fmt.Printf("The type is %T\n", a)
	var b int = 20
	b = 30
	fmt.Println("b:", b)
	// a = 20 not possible
	// value of const should be known on compile time not at runtime
	// const b = callMethod() not possible
	const str = "string" // type of string constant is untyped
	fmt.Println("str:", str)
	fmt.Printf("The type is %T\n", str)

	// every variable type should be explicitly mentioned
	// in variable it recognizes from the value
	// in constant except string others are recognized

	var defaultName = "Sam"         //allowed
	type myString string            // making myString alias for string
	var customName myString = "Sam" //allowed

	fmt.Println(defaultName, customName)
	// customName = defaultName  not allowed
	// because though myString is string, Go's strong type checking tells both are different

	const trueConst = true
	type myBool bool
	var defaultBool = trueConst       //allowed
	var customBool myBool = trueConst //allowed

	fmt.Println(defaultBool, customBool)
	//defaultBool = customBool not allowed
	// same as string constants

	// constants are untyped
	// when used as initialization for other variables, type should be specified along with it
	var intVar = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
	fmt.Printf("Type of intVar is %T\n", intVar)
	fmt.Printf("Type of int32Var is %T\n", int32Var)
	fmt.Printf("Type of float64Var is %T\n", float64Var)
	fmt.Printf("Type of complex64Var is %T\n", complex64Var)

	var z = 5.6 / 3
	fmt.Println("Division")
	fmt.Println("Value of z is ", z)
	fmt.Printf("The type of z is %T", z)

	z = 5.6 * 3
	fmt.Println("Multiplication")
	fmt.Println("Value of z is ", z)
	fmt.Printf("The type of z is %T", z)

	z = 5.6 + 3
	fmt.Println("Addition")
	fmt.Println("Value of z is ", z)
	fmt.Printf("The type of z is %T", z)

	z = 5.6 - 3
	fmt.Println("Substraction")
	fmt.Println("Value of z is ", z)
	fmt.Printf("The type of z is %T", z)
}
