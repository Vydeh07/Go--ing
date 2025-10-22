package main

import "fmt"

func main() {

	// sequence := adder()

	// fmt.Println(sequence()) // Output: 1
	// fmt.Println(sequence()) // Output: 2
	// fmt.Println(sequence()) // Output: 3
	// fmt.Println(sequence()) // Output: 4
	// fmt.Println(sequence()) // Output: 5dddddddd

	// sequence2 := adder()
	// fmt.Println(sequence2())/// again 1 tfff

	subtracter := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	//using the closure subtracter
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(4))

}

func adder() func() int {
	i := 0
	fmt.Println("Previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
