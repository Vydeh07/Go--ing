package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("This will not be printed because of os.Exit")
	fmt.Println("Starting the main function")

	os.Exit(1)

	fmt.Println("End of main function")
}
