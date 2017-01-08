package main

import (
	"math"
	"strconv"
)

type FizzBuzzType interface {
	result(number float64) string
}

type Fizz struct {
	next FizzBuzzType
}

func (fizzType *Fizz) result(number float64) string {
	if 0 == math.Mod(number, 3) {
		return "Fizz"
	}
	return fizzType.next.result(number)
}

type Buzz struct {
	next FizzBuzzType
}

func (buzzType *Buzz) result(number float64) string {
	if 0 == math.Mod(number, 5) {
		return "Buzz"
	}
	return buzzType.next.result(number)
}

type DefaultType struct{}

func (defaultType *DefaultType) result(number float64) string {

	return strconv.FormatFloat(number, 'f', -1, 64)
}

var fizzBuzzTypes = &Fizz{&Buzz{&DefaultType{}}}

func FizzBuzz(number float64) string {
	return fizzBuzzTypes.result(number)
}
