package main

import "fmt"

func main() {
	//simple iteration over a range
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	//iterate over collection
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
