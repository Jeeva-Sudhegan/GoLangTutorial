package package12

/*
	to export the function write the function name with caps at start
	eg. Area, Perimeter, etc.
*/

func Area(length, breadth float64) (area float64) {
	area = length * breadth
	return
}

func Perimeter(length, breadth float64) (perimeter float64) {
	perimeter = 2 * (length + breadth)
	return
}
