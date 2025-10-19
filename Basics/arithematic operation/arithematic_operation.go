package main

import "fmt"

func main() {
	var a, b int = 10, 3
	var result int
	result = a + b
	fmt.Println("Addition: ", result)
	result = a - b
	fmt.Println("Subtraction:", result)

	const p float64 = 22 / 7
	fmt.Println(p) // we get 3 and if we give 7 as 7.0 we get 3.142857142857143

	//overflow with signed integer
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt) // -ve

	//overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615
	fmt.Println(maxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt) // 0

}
