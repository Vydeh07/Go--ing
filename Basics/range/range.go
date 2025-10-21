package main

import "fmt"

func main() {
	message := "HHH World"

	for i, v := range message {
		//fmt.Println(i, v)
		fmt.Printf("index: %d, Rune: %c\n", i, v)
	}
}
