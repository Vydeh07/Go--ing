package main

import "fmt"

func main() {
	// func functionName(param1 type1, param2 type2, params ... typeN) returnType {
	// 	// function body
	// 	return returnValue
	// }

	// fmt.Println("Sum of 1,2,3:", sum(1, 2, 3, 4, 5, 6, 7, 78))
	statement, total := sum("The total sum is:", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(statement, total)

}
func sum(returnString string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return returnString, total
}
