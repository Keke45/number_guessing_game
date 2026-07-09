package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const MaxNumber = 100
	secretNumber := rand.Intn(MaxNumber) + 1
	attempts := 0
	var guess int
	fmt.Println("Guess a number between 1 and 100")

	for {

		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		attempts++

		if guess < secretNumber {
			fmt.Println("Too low!")
		} else if guess > secretNumber {
			fmt.Println("Too high!")
		} else {
			fmt.Printf("Congratulations! You guessed the number %d in %d attempts.\n", secretNumber, attempts)
			break
		}
	}
}
