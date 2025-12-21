package main

import (
	"math/rand/v2"
	"strconv"
	"strings"
)

// GenerateRandomNumber returns a zero-padded random numeric string
// with exactly the specified length.
//
// The function generates a random number whose maximum value is based
// on the given length and converts it to a string. If the generated
// number has fewer digits than the requested length, it is left-padded
// with zeros to ensure the returned string always has the exact length.
//
// Example:
//
//	GenerateRandomNumber(4) -> "0382"
//	GenerateRandomNumber(6) -> "902145"
//
// Note:
//   - The returned value is a string, not an integer.
//   - Leading zeros may be present.
//   - The randomness depends on the global rand source; callers should
//     seed math/rand if deterministic behavior is not desired.
//
// Parameters:
//
//	length - the desired number of digits in the returned string.
//
// Returns:
//
//	A string containing a random numeric value of exactly `length` digits.
func GenerateRandomNumber(length int) string {
	if length <= 0 {
		return ""
	}

	// get the max number of digits for random
	max := 10
	for i := 1; i < length; i++ {
		max *= 10
	}

	// range 0 to max-1 [0, max)
	// convert to string
	num := strconv.Itoa(rand.IntN(max))

	// check if the generated number is equals to length
	// if not equal to length then add 0 to the string
	if diff := length - len(num); diff != 0 {
		num = strings.Repeat("0", diff) + num
	}

	return num
}
