package main

import "fmt"

func main() {

	fmt.Println(factorial(5))

	fmt.Println(sumOfDigits(9))
	fmt.Println(sumOfDigits(155))
	fmt.Println(sumOfDigits(15))

}
func factorial(n int) int {
	//Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}
	//Recursive case: n! = n * (n-1)!
	return n * factorial(n-1)

}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}
