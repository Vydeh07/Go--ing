package main

import "fmt"

func main() {
	process()
	fmt.Println("returned from process")
}

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}

	}()
	fmt.Println("Starting process")
	panic("Something went wrong!")
	fmt.Println("Process completed")
}
