package main

import (
	"fmt"
	"unicode/utf8"
)

/*
	slice of bytes
	enclosed in ""
	unicode compliant and utf-8 encoded
	accessing each element in string is actually we are getting bytes not characters
	fmt.Printf("% x",x) means print each individual with spaces
	fmt.Printf("%q",x) means print x that escapes if non-printable character present
	fmt.Printf("%+q",x) means print x that escapes if non-printable and non ASCII characters are present
*/

func main() {
	name := "Jeeva"       // string declaration
	runes := []rune(name) // rune is alias for int32. represents unicode code point. doesn't matter how many bytes needed for one character, it is considered as one
	fmt.Println("Hello", name)
	fmt.Println("Rune:", runes)
	for i := 0; i < len(name); i++ {
		fmt.Print(name[i])        // display unicode index of the character
		fmt.Printf("%c", name[i]) // to display character
	}
	// string slices
	fmt.Println()
	fmt.Println("To understand how string works in GO")
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	var sliceSample = []string{"\xbd", "\xb2", "\x3d", "\xbc", "\x20", "\xe2", "\x8c", "\x98"}

	fmt.Println("Println:", sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x\t", sample[i]) // prints hex value
		fmt.Printf("%q\t", sample[i]) // prints symbols
		fmt.Printf("%+q", sample[i])  // prints unicode for symbols
		fmt.Println()
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)

	fmt.Println("Println:", sliceSample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sliceSample); i++ {
		fmt.Printf("%x\t", sliceSample[i]) // prints the hex value
		fmt.Printf("%q\t", sliceSample[i]) // prints hex string
		fmt.Printf("%+q", sliceSample[i])  // same as above
		fmt.Println()
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sliceSample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sliceSample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sliceSample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sliceSample)

	fmt.Println("sample and sliceSample is different")

	/*
		Go source code uses utf-8
		utf-8 is created when source code is written
	*/

	fmt.Println("Understanding the encoding")

	const placeOfInterest = `⌘` // utf-8 code is replaced into the placeOfInterest

	// string literals are utf-8
	// string values are arbitrary bytes
	// string contains bytes but when constructed from literals, those bytes are utf-8

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i]) // utf-8 encoding of the bytes
	}
	fmt.Printf("\n")

	/*
		Go source code is always UTF-8.
		A string holds arbitrary bytes.
		A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
		Those sequences represent Unicode code points, called runes.
		No guarantee is made in Go that characters in strings are normalized.
		Go treats UTF-8 specially when using for range loop
	*/
	fmt.Println("Japanese letter")
	const nihongo = "日本語"
	var runed = []rune(nihongo)
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	// if used normal for loop, it will print the combination bytes for the above characters
	for index, runeValue := range runed {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	// if used rune, it will be like normal array of bytes with each character in respective indices
	var strRune = string(runed)
	fmt.Println("length of nihongo:", len(nihongo))                           // it also prints the required code points of the bytes
	fmt.Println("Actual length of nihongo:", utf8.RuneCountInString(nihongo)) // prints actual length since each characters code points are considered
	fmt.Println("Build from rune:", strRune)

	// Strings are immutable
	// to mutate string, convert it to slice of rune and mutate and convert back to string
	// name[0] = "j" throws compiler error
	var nameRune = []rune(name)
	nameRune[0] = 'j' // no double quotes since it is a character
	name = string(nameRune)
	fmt.Println("After mutating", name)
}
