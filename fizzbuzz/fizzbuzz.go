package fizzbuzz

import "strconv"

// Fizzbuzz : fizzbuzz
func Fizzbuzz(s string) string {
	n, _ := strconv.Atoi(s)
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return s
}
