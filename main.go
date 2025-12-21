package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	attempts := 0

	fmt.Println("Welcome to the Code Cracker!")
	fmt.Println()
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (3 digits)")
	fmt.Println("2. Medium (4 digits)")
	fmt.Println("3. Hard (5 digits)")
	fmt.Println()
	num := SetDifficulty()
	target := GenerateRandomNumber(num)
	start := time.Now()

	for {
		fmt.Print("Enter your guess: ")
		scanner.Scan()
		guess := scanner.Text()
		hit, bullsEye := Check(guess, target)
		attempts += 1

		if bullsEye != num {
			fmt.Printf("Incorrect! You got %d bullseyes and %d hits.\n\n", bullsEye, hit)
			continue
		}
		break
	}
	elapsed := time.Since(start)
	fmt.Printf("Congratulations! You guessed the correct number in %d attempts and %.2f seconds.\n", attempts, elapsed.Seconds())
}
