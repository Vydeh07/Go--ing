package main

import "fmt"

func main() {
	//panic(interface{})
	process(10)
	process(-5)

}

func process(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before panic")
		panic("Negative input not allowed ")
		fmt.Println("After Panic")

		defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input:", input)
}
