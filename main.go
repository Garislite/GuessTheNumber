package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secretNumber := rand.Intn(100) + 1
	var guess int
	attempts := 0

	fmt.Println("GUESS THE NUMBER")
	fmt.Println("=================")
	fmt.Println("I picked a number from 1 to 100")

	for {
		fmt.Print("\nYour guess: ")
		fmt.Scan(&guess)
		attempts++

		if guess < secretNumber {
			fmt.Println("Too low")
		} else if guess > secretNumber {
			fmt.Println("Too high")
		} else {
			fmt.Printf("\nCorrect! The number was %d\n", secretNumber)
			fmt.Printf("Attempts: %d\n", attempts)
			break
		}
	}
}
