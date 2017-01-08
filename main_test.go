package main

import "testing"

var testFizzStruct = []struct {
	input    float64
	expected string
}{
	{1, "1"},
	{2, "2"},
	{3, "Fizz"},
	{6, "Fizz"},
}
var testBuzzStruct = []struct {
	input    float64
	expected string
}{
	{1, "1"},
	{2, "2"},
	{5, "Buzz"},
	{10, "Buzz"},
	{11, "11"},
}
var testFizzBuzzStruct = []struct {
	input    float64
	expected string
}{
	{1, "1"},
	{2, "2"},
	{15, "FizzBuzz"},
	{30, "FizzBuzz"},
	{11, "11"},
}

func TestShouldResponseFizzInCaseOfMultipleOfThree(t *testing.T) {
	for _, structure := range testFizzStruct {
		stringObtained := GoFizzBuzz(structure.input)

		if stringObtained != structure.expected {
			t.Fail()
		}
	}
}

func TestShouldResponseBuzzInCaseOfMultipleOfFive(t *testing.T) {
	for _, structure := range testBuzzStruct {
		stringObtained := GoFizzBuzz(structure.input)

		if stringObtained != structure.expected {
			t.Fail()
		}
	}
}

func TestShouldResponseFizzBuzzInCaseOfMultipleOfEitherThreeAndFive(t *testing.T) {
	for _, structure := range testFizzBuzzStruct {
		stringObtained := GoFizzBuzz(structure.input)

		if stringObtained != structure.expected {
			t.Fail()
		}
	}
}
