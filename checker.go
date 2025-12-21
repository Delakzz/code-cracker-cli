package main

import "strings"

func Check(guess, target string) (int, int, bool) {
	hit, bullsEye := 0, 0

	// check if the user guessed the target random number correctly
	if x := strings.Compare(guess, target); x == 0 {
		return len(target), 0, true
	}

	// create maps for easier checking
	m := make(map[string]int)

	// count bullsEye
	for i := 0; i < len(guess); i++ {
		x, y := string(guess[i]), string(target[i])

		// check if it is hit
		if x == y {
			bullsEye += 1
		}

		// add to maps
		m[y] = i
	}

	// count hit
	for i := 0; i < len(guess); i++ {
		if num := m[string(guess[i])]; num != 0 && num != i+1 {
			hit += 1
		}
	}

	return hit, bullsEye, false
}
