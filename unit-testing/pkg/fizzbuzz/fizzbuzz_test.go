package fizzbuzz

import (
	"strconv"
	"testing"
)

// A test to ensure the result is Fizz for all values
// divisible by fizzAt and not buzzAt
func TestFizzAt(t *testing.T) {
	total := int64(100)
	fizzAt := int64(3)
	buzzAt := int64(5)

	result := FizzBuzz(total, fizzAt, buzzAt)

	for i := int64(0); i < total; i += fizzAt {
		if i != 0 && i%fizzAt == 0 && !(i%buzzAt == 0) {
			if result[i-1] != "Fizz" {
				t.Fail()
			}
		}
	}
}

// A test to ensure the result is Buzz for all values
// divisible by buzzAt and not fizzAt
func TestBuzzAt(t *testing.T) {
	total := int64(100)
	fizzAt := int64(3)
	buzzAt := int64(5)

	result := FizzBuzz(total, fizzAt, buzzAt)

	for i := int64(0); i < total; i += buzzAt {
		if i != 0 && i%buzzAt == 0 && !(i%fizzAt == 0) {
			if result[i-1] != "Buzz" {
				t.Fail()
			}
		}
	}
}

// A test to ensure the result is FizzBuzz for all values
// divisible by fizzAt and buzzAt
func TestFizzBuzzAt(t *testing.T) {
	total := int64(100)
	fizzAt := int64(3)
	buzzAt := int64(5)

	result := FizzBuzz(total, fizzAt, buzzAt)

	for i := int64(0); i < total; i += (fizzAt * buzzAt) {
		if i != 0 && i%buzzAt == 0 && i%fizzAt == 0 {
			if result[i-1] != "FizzBuzz" {
				t.Fail()
			}
		}
	}
}

// A test to ensure the result is a string containing Fizz, Buzz or FizzBuzz
// for all values divisible by fizzAt OR buzzAt
func TestFizzBuzzOutput(t *testing.T) {
	total := int64(100)
	fizzAt := int64(3)
	buzzAt := int64(5)

	result := FizzBuzz(total, fizzAt, buzzAt)

	for i := int64(0); i < total; i++ {
		if i != 0 && (i%buzzAt == 0 || i%fizzAt == 0) {
			if !((result[i-1] == "Fizz") ||
				(result[i-1] == "Buzz") ||
				(result[i-1] == "FizzBuzz")) {
				t.Fail()
			}
		}
	}
}

// A test to ensure the result is number and that
//number is equal to the value FizzBuzzed
func TestNumberOutput(t *testing.T) {
	total := int64(100)
	fizzAt := int64(3)
	buzzAt := int64(5)

	result := FizzBuzz(total, fizzAt, buzzAt)

	for i := int64(0); i < total; i++ {
		if i != 0 && !(i%buzzAt == 0) && !(i%fizzAt == 0) {
			val, err := strconv.Atoi(result[i-1])
			if err != nil || int64(val) != i {
				t.Fail()
			}
		}
	}
}

// A test to ensure that a negative total returns a string array of len 1
// containing an error message
func TestNegativeTotal(t *testing.T) {
	total := int64(-100)
	fizzAt := int64(3)
	buzzAt := int64(5)

	errorMsg := "error: total less than 0, invalid integer range"

	result := FizzBuzz(total, fizzAt, buzzAt)
	if len(result) != 1 || result[0] != errorMsg {
		t.Fail()
	}

}
