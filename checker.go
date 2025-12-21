package main

import (
	"errors"
	"strconv"
)

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
func Check(guess, target string) (int, int, error) {
	hit, bullsEye := 0, 0

	// validate length
	if len(guess) != len(target) {
		return 0, 0, errors.New(
			"Invalid input. It must be a " + strconv.Itoa(len(target)) + "-digit number.\n",
		)
	}

	// maps to count remaining digits (excluding bullsEye)
	guessCount := make(map[byte]int)
	targetCount := make(map[byte]int)

	// count bullsEye
	for i := 0; i < len(guess); i++ {
		if guess[i] == target[i] {
			bullsEye++
		} else {
			// add only non-bullsEye digits
			guessCount[guess[i]]++
			targetCount[target[i]]++
		}
	}

	// count hits
	for digit, gCount := range guessCount {
		if tCount, ok := targetCount[digit]; ok {
			// hit is the minimum of occurrences
			// gCount and tCount is the occurences of a digit
			// so if gCount is 1 and tCount is 2, the hit will only count it as 1
			// if gCount is 2 and tCount is 1, it will also count as 1
			// this is to avoid the duplications of hit
			hit += min(gCount, tCount)
		}
	}

	return hit, bullsEye, nil
}
