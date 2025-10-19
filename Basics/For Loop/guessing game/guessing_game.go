package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	//generate a random number between 1 and 100

	target := random.Intn(100) + 1

	//Welcome Message
	fmt.Println("Welcome to guessing  game")
	fmt.Println("I have chosen a number between 1 to 100,,,  Guess it!")

	var guess int
	for {
		fmt.Println("Enter your guess")
		fmt.Scanln(&guess)

		///Check if the guess is correct
		if guess == target {
			fmt.Println("congrats! you guessed it right")
			break
		} else if guess < target {
			fmt.Println("Too low! Try again.")

		} else {
			fmt.Println("Too high! Try again.")
		}
	}

}
