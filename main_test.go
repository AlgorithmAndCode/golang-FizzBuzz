package main

import "testing"

var testStruct = []struct {
	input    float64
	expected string
}{
	{1, "1"},
	{2, "2"},
	{3, "Fizz"},
	{6, "Fizz"},
}

func TestShouldResponseFizzInCaseOfMultipleOfThree(t *testing.T) {
	for _, structure := range testStruct {
		stringObtained := FizzBuzz(structure.input)

		if stringObtained != structure.expected {
			t.Fail()
		}
	}
}
