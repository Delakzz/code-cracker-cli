package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	attempts := 0
	target := GenerateRandomNumber(4)
	guessed := false

	fmt.Println("Welcome to the Code Cracker!")
	fmt.Println("I'm thinking of a 4-digit number.")
	fmt.Println("You have 5 chances to guess the correct number.")
	fmt.Println()

	for attempts < 5 {
		fmt.Print("Enter your guess: ")
		scanner.Scan()
		guess := scanner.Text()
		hit, bullsEye := Check(guess, target)
		attempts += 1

		if bullsEye == 4 {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempts)
			guessed = true
			break
		} else {
			fmt.Printf("Incorrect! You got %d bullseyes and %d hits.\n\n", bullsEye, hit)
		}

	}

	if !guessed {
		fmt.Printf("The target number is %s!!\n", target)
	}
}
