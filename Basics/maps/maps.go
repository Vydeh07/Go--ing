package main

import (
	"fmt"
	"maps"
)

func main() {
	// var m map [key] valueType

	myMap := make(map[string]int)

	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18

	fmt.Println(myMap)
	fmt.Println(myMap["key1"])

	delete(myMap, "key1")
	fmt.Println(myMap)

	myMap["code0"] = 18
	myMap["code1"] = 19
	myMap["code2"] = 20

	//clear(myMap)
	fmt.Println(myMap)

	value, unknownvalue := myMap["key1"]
	fmt.Println(value)
	fmt.Println("is a value associated with key1:", unknownvalue)

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	if maps.Equal(myMap, myMap2) {
		fmt.Println("myMap and myMap2 are equal")

	}

	for k, v := range myMap2 {
		fmt.Println(k, v)
	}
	for _, v := range myMap2 {
		fmt.Println(v)
	}

	var myMap4 map[string]string
	if myMap4 == nil {
		fmt.Println("The map is initialized to nil")
	} else {
		fmt.Println("The map is not  initialized")
	}
	myMap4 = make(map[string]string)
	myMap4["key"] = "Value"
	fmt.Println(myMap4)

	fmt.Println("myMap4 length is ", len(myMap4))

	myMap5 := make(map[string]map[string]string)

	myMap5["map1"] = myMap4
	fmt.Println(myMap5)

}
