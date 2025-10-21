package main

import "fmt"

func main() {
	//func <name>(parameter list) return type}{
	//code nlock
	//return value
	//}
	// 	sum := add(1, 2)
	// 	fmt.Println(sum)

	// }

	//

	//passing a function as an argument
	result := applyOperation(5, 3, add)
	fmt.Println("5 + 3", result)

	//return and using a func
	muliplyby2 := createMultiplier(2)
	fmt.Println("6 * 2", muliplyby2(6))
}
func add(a, b int) int {
	return a + b
}

//function that takes a function as an argument

func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

//function that returns a function

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
