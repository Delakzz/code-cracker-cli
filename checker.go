package main

import "strings"

// Check compares the player's guess with the target number in the game
// and returns the number of hits and bulls-eye
//
// Parameters:
//   - guess  : The player's guessed number as a string.
//   - target : The target number to be guessed as a string.
//
// Returns:
//   - hit       : The count of digits that exist in the target but are in the wrong position.
//   - bullsEye  : The count of digits that are correct and in the correct position.
//
// Behavior:
//   - If the guess is exactly the same as the target, the function returns the length
//     of the target as bulls-eye, and 0 hits.
//   - Otherwise, it counts bulls-eye first (correct digits in the correct position),
//     then counts hits (correct digits in the wrong position) using a map for efficiency.
func Check(guess, target string) (int, int) {
	hit, bullsEye := 0, 0

	// return immediately if guess exactly matches target
	if x := strings.Compare(guess, target); x == 0 {
		return len(target), 0
	}

	// create map for easier checking
	m := make(map[string]int)

	// count bullsEye
	for i := 0; i < len(guess); i++ {
		x, y := string(guess[i]), string(target[i])

		// check if it is hit
		if x == y {
			bullsEye += 1
		}

		// add to maps
		// +1 because 0 is used to check if the number exists
		m[y] = i + 1
	}

	// count hit
	for i := 0; i < len(guess); i++ {
		// num != 0: exists
		// num != i+1: not in the same position
		if num := m[string(guess[i])]; num != 0 && num != i+1 {
			hit += 1
		}
	}

	return hit, bullsEye
}
