package main

import "fmt"

func main() {
	// age := 25
	// if age >= 18 {
	// 	fmt.Println("You are an adult")
	// }

	temperature := 25
	if temperature > 30 {
		fmt.Println("It's a hot day")
	} else {
		fmt.Println("Its cool outside")
	}

	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}
}
