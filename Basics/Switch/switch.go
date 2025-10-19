package main

import "fmt"

func main() {
	// switch expression {
	// case value1:
	// 	//fallthrough // check next case->fallthrough no need of break
	// case value2:
	// 	//
	// default:
	// 	//
	// }

	// fruit := "apple"

	// switch fruit {
	// case "apple":
	// 	fmt.Println("Its a apple")
	// case "banana":
	// 	fmt.Println("Its a banana")
	// default:
	// 	fmt.Println("Unknown fruit")
	// }

	// // multiple condition
	// day := "monday"

	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday")
	// case "Sunday":
	// 	fmt.Println("It's a weekend")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// number := 15

	// switch {
	// case number < 10:
	// 	fmt.Println("Number is less than 10")
	// case number >= 10 && number <= 20:
	// 	fmt.Println("Number is between 10 and 20")
	// default:
	// 	fmt.Println("Number is greater than 20")
	// }

	// //fallthrough

	// num := 2
	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("Equal to 2")
	// default:
	// 	fmt.Println("Less than 1")
	// }

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
		//cant use fallthrough in type switch
	case float64:
		fmt.Println("It's a float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown type")
	}
}
