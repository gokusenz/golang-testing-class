package fizzbuzz

import "testing"

// Define a function ShareWith(string) string.

func TestSendOneShouldBeSameInput(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"1", "1"},
		{"2", "2"},
	}
	for _, test := range tests {
		if observed := Fizzbuzz(test.name); observed != test.expected {
			t.Fatalf("it should be (%s) = %s", observed, test.expected)
		}
	}
}

func TestSendOneShouldBeFizz(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"3", "Fizz"},
		{"9", "Fizz"},
	}
	for _, test := range tests {
		if observed := Fizzbuzz(test.name); observed != test.expected {
			t.Fatalf("it should be (%s) = %s", observed, test.expected)
		}
	}
}

func TestSendOneShouldBeBuzz(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"10", "Buzz"},
		{"5", "Buzz"},
	}
	for _, test := range tests {
		if observed := Fizzbuzz(test.name); observed != test.expected {
			t.Fatalf("it should be (%s) = %s", observed, test.expected)
		}
	}
}

func TestSendOneShouldBeFizzBuzz(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"15", "FizzBuzz"},
		{"30", "FizzBuzz"},
	}
	for _, test := range tests {
		if observed := Fizzbuzz(test.name); observed != test.expected {
			t.Fatalf("it should be (%s) = %s", observed, test.expected)
		}
	}
}
