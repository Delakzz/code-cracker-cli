package main

import (
	"math/rand/v2"
	"strconv"
	"strings"
)

func GenerateRandomNumber(length int) string {
	// get the max number of digits for random
	base := 10
	for i := 1; i < length-1; i++ {
		base *= 10
	}

	// range 0 to base * 9
	// convert to string
	num := strconv.Itoa(rand.IntN(base * 9))

	// check if the generated number is equals to length
	// if not equal to length then add 0 to the string
	if diff := length - len(num); diff != 0 {
		x := "0"
		res := strings.Repeat(x, diff)
		res += num
		num = res
	}

	// return
	return num
}
