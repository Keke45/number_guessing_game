package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const MaxNumber = 100
const MaxAttempts = 5

func generateSecretNumber(max int) int {
	return rand.Intn(max) + 1
}

func getGuess(scanner *bufio.Scanner, max int) (int, error) {
	for {
		fmt.Print("Enter your guess: ")

		if !scanner.Scan() {
			return 0, errors.New("failed to read input")
		}
		text := scanner.Text()
		guess, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Please enter a valid whole number.")
			continue
		}
		if guess < 1 || guess > max {
			fmt.Printf("Please enter a number between 1 and %d\n", max)
			continue
		}
		return guess, nil
	}
}

func checkGuess(secret, guess int) string {
	switch {
	case guess < secret:
		return "Too low!"
	case guess > secret:
		return "Too high!"
	default:
		return "Correct!"
	}
}

func main() {

	secret := generateSecretNumber(MaxNumber)
	scanner := bufio.NewScanner(os.Stdin)
	attempts := 0
	won := false
	fmt.Printf("Guess a number between 1 and %d\n", MaxNumber)

	for attempts < MaxAttempts {
		guess, err := getGuess(scanner, MaxNumber)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		attempts++

		result := checkGuess(secret, guess)
		fmt.Println(result)

		if result == "Correct!" {
			fmt.Printf("You guessed the number in %d attempts!\n", attempts)
			won = true
			break
		}
		fmt.Printf("Attempts remaining: %d\n", MaxAttempts-attempts)

	}
	if !won {
		fmt.Printf("Game over! The number was %d.\n", secret)
	}
}
