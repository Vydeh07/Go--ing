package main

import "fmt"

func main() {
	process()
}

func process() {
	defer fmt.Println("Deferred: This will be printed last")
	defer fmt.Println("Deferred: This will be printed second last")
	defer fmt.Println("Deferred: This will be printed third last")
	fmt.Println("Normal: This will be printed first")
}
