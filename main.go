package main

import (
	"math"
	"strconv"
)

func FizzBuzz(number float64) string {
	if 0 == math.Mod(number, 3) {
		return "Fizz"
	}
	return strconv.FormatFloat(number, 'f', -1, 64)
}
