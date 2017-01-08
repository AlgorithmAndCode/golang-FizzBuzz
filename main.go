package fizzbuzz

import (
	"math"
	"strconv"
)

type Resolver interface {
	Resolve(n float64) string
}

type FType struct {
	next Resolver
}

func (ft *FType) Resolve(n float64) string {
	if 0 == math.Mod(n, 3) {
		return "Fizz"
	}
	return ft.next.Resolve(n)
}

type BType struct {
	next Resolver
}

func (bt *BType) Resolve(n float64) string {
	if 0 == math.Mod(n, 5) {
		return "Buzz"
	}
	return bt.next.Resolve(n)
}

type FBType struct {
	next Resolver
}

func (fbt *FBType) Resolve(n float64) string {
	if 0 == math.Mod(n, 3) && 0 == math.Mod(n, 5) {
		return "FizzBuzz"
	}
	return fbt.next.Resolve(n)
}

type DefaultType struct{}

func (dt *DefaultType) Resolve(n float64) string {

	return strconv.FormatFloat(n, 'f', -1, 64)
}

var resolvers = &FBType{&FType{&BType{&DefaultType{}}}}

func DoFizzBuzz(n float64) string {
	return resolvers.Resolve(n)
}
