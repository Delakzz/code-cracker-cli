package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// SetDifficulty prompts the user to select a difficulty level for the game
// and returns the corresponding number of digits for the random number.
//
// Difficulty Levels:
//
//	1 -> Easy   : 3 digits
//	2 -> Medium : 4 digits
//	3 -> Hard   : 5 digits
//
// Returns:
//
//	int : Number of digits based on the selected difficulty level.
//
// Behavior:
//   - Continuously asks the user for input until a valid choice (1, 2, or 3) is entered.
//   - Prints confirmation of the selected difficulty level.
func SetDifficulty() int {
	var num int
	scanner := bufio.NewScanner(os.Stdin)

	// map for getting the difficulty and digits
	m := map[int]string{1: "Easy", 2: "Medium", 3: "Hard"}
	n := map[string]int{"easy": 3, "Medium": 4, "Hard": 5}

	for {
		fmt.Print("Enter your choice:")
		scanner.Scan()
		choice := scanner.Text()
		num, _ := strconv.Atoi(choice)

		// check if the user input is valid or not
		if num < 1 || num > 3 {
			fmt.Printf("\nInvalid input. Please pick again.\n")
			continue
		}
		break
	}

	fmt.Printf("\nGreat! You have selected the %s difficulty level.\n", m[num])
	fmt.Println("Let's start the game!")
	return n[m[num]]
}
