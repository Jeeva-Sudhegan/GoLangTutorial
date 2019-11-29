package main

import "fmt"

// same as map datastructure
// these are reference types like slice
// cannot be compared using ==
// == is used only to check if the map is 'nil' or not
// to check both maps same or not, check individual items
func main() {
	mapObj := make(map[string]string) // to create map => make(map[<key type>]<value type>)
	// declaring
	var anotherMap map[string]string
	if anotherMap == nil {
		anotherMap = make(map[string]string)
	}
	anotherMap["hello"] = "world"
	// defining on declaring map
	var thirdMap map[string]string = map[string]string{
		"hello": "world",
	}
	mapObj["hello"] = "world" // to add items to map => map[key] = value
	world := mapObj["hello"]  // to get value for given key
	// if key is not present, return the default value for the type of the value
	check, ok := mapObj["world"] // there is no key, it returns ''. ok holds whether the key is present in map or not
	fmt.Println(mapObj)
	fmt.Println(thirdMap)
	fmt.Println("mapObj['hello']:", world)
	if !ok {
		fmt.Println("'world' is not in mapObj")
	}
	fmt.Println("mapObj['world']:", check)
	for key, value := range mapObj { // for iterating map. Order cannot be guranteed when using range
		fmt.Printf("mapObj[%s] = %s\n", key, value)
	}
	delete(mapObj, "hello") // to delete key from map => delete(<map_name>,<key>)
	fmt.Println("After deleted mapObj:", mapObj)
	length := len(mapObj) // length of map
	fmt.Println("Length of mapObj", length)
	equal := isEqual(anotherMap, thirdMap)
	if equal {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
}

// checking if two maps are equal
func isEqual(firstMap, secondMap map[string]string) bool {
	if len(firstMap) == len(secondMap) {
		for firstKey, firstValue := range firstMap {
			secondValue, ok := secondMap[firstKey]
			if !ok || secondValue != firstValue {
				return false
			}
		}
		return true
	}
	return false
}
