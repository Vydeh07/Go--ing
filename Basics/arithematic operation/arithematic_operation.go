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

}
