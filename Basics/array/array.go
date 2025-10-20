package main

import "fmt"

func main() {
	var numbers [5]int

	fmt.Println("Initial numbers array:", numbers)

	numbers[1] = 20
	fmt.Println("After updating numbers array:", numbers)

	fruits := [4]string{"apple", "banana", "mango", "grapes"}
	fmt.Println("Fruits array:", fruits)

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray
	copiedArray[0] = 100

	fmt.Println("Original Array:", originalArray)
	fmt.Println("Copied Array:", copiedArray)

	// Loop through numbers array using traditional for loop
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index", i, "is", numbers[i])
	}

	// Loop through numbers array using range
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Loop through numbers array ignoring index
	for _, v := range numbers {
		fmt.Printf("Value: %d\n", v)
	}

	fmt.Println("The length of numbers array is", len(numbers))

	// Comparing arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	fmt.Println("Are array1 and array2 equal?", array1 == array2) // true

	// 2D array (Matrix)
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Print the matrix in a readable format
	fmt.Println("Matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
