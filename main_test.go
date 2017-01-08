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
var testAllTypesStruct = []struct {
	input    float64
	expected string
}{
	{1, "1"},
	{2, "2"},
	{3, "Fizz"},
	{4, "4"},
	{5, "Buzz"},
	{6, "Fizz"},
	{7, "7"},
	{8, "8"},
	{9, "Fizz"},
	{10, "Buzz"},
	{11, "11"},
	{12, "Fizz"},
	{15, "FizzBuzz"},
	{30, "FizzBuzz"},
	{34, "34"},
	{50, "Buzz"},
	{90, "FizzBuzz"},
	{100, "Buzz"},
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

func TestShouldResponseAppropriateStringInAnyPosibleCase(t *testing.T) {
	for _, structure := range testAllTypesStruct {
		stringObtained := GoFizzBuzz(structure.input)

		if stringObtained != structure.expected {
			t.Fail()
		}
	}
}

