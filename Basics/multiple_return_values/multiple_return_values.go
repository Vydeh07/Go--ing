package main

import (
	"errors"
	"fmt"
)

func main() {
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d Remainder :%d\n", q, r)

	result, err := compare(2, 2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func compare(a, b int) (string, error) {
	if a > b {
		return "A is greater than B", nil
	} else if b > a {
		return "B is greater than A", nil
	} else {
		return "", errors.New("unable to compare which is greater")
	}
}
