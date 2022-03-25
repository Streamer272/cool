package main

import (
	"fmt"
	"strconv"
)

type Int int
type IntArray []Int

func NewInt(i int) Int {
	return Int(i)
}

func NewIntFromString(s string) Int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return Int(i)
}

func (i *Int) Int() int {
	return int(*i)
}

func (i *Int) ToString() string {
	return strconv.Itoa(int(*i))
}

func (i *Int) Increment() {
	*i++
}

func (i *Int) IncrementBy(by int) {
	*i = Int(int(*i) + by)
}

func (i *Int) Decrement() {
	*i--
}

func (i *Int) DecrementBy(by int) {
	*i = Int(int(*i) - by)
}

func (i *Int) MustBePositive() {
	if *i < 0 {
		panic("Int must be positive")
	}
}

func (i *Int) MustBeNegative() {
	if *i > 0 {
		panic("Int must be negative")
	}
}

func (i *Int) MustBeZero() {
	if *i != 0 {
		panic("Int must be zero")
	}
}

func (i *Int) MustEqual(other *Int) {
	if *i != *other {
		panic(fmt.Sprintf("Int (%d) must equal %d", *i, *other))
	}
}

func (i *Int) MustBeEqualTo(other int) {
	if int(*i) != other {
		panic(fmt.Sprintf("Int (%d) must equal %d", *i, other))
	}
}
