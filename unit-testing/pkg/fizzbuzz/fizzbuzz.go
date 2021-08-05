package fizzbuzz

import (
	"strconv"
)

// FizzBuzz performs a FizzBuzz operation over a range of integers
//
// Given a range of integers:
// - Return "Fizz" if the integer is divisible by the `fizzAt` value.
// - Return "Buzz" if the integer is divisible by the `buzzAt` value.
// - Return "FizzBuzz" if the integer is divisible by both the `fizzAt` and
//   `buzzAt` values.
// - Return the original number if is is not divisible by either the `fizzAt` or
//   the `buzzAt` values.
func FizzBuzz(total, fizzAt, buzzAt int64) []string {
	// A check to prevent a negative total giving us a makeslice: len out of range error
	// Returns string array of lenght 1 with an error message
	if total < 0 {
		result := make([]string, 1)
		result[0] = "error: total less than 0, invalid integer range"
		return result
	}

	result := make([]string, total)

	for i := int64(1); i <= total; i++ {
		if !(i%fizzAt == 0) && !(i%buzzAt == 0) {
			result[i-1] = strconv.FormatInt(i, 10)
			continue
		}

		if i%fizzAt == 0 {
			result[i-1] = "Fizz"
		}

		if i%buzzAt == 0 {
			result[i-1] += "Buzz"
		}
	}

	return result
}
