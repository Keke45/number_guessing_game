package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	const MaxNumber = 100
	secretNumber := rand.Intn(MaxNumber) + 1
	attempts := 0
	var guess int
	fmt.Printf("Guess a number between 1 and %d\n", MaxNumber)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter your guess: ")
		scanner.Scan()
		text := scanner.Text()

		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		guess = num

		if guess < 1 || guess > MaxNumber {
			fmt.Printf("Please enter a number between 1 and %d.\n", MaxNumber)
			continue
		}
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
