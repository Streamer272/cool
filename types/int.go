package types

import (
	"fmt"
	"strconv"
)

type Int struct {
	int
}
type IntArray []Int

func NewInt(i int) Int {
	return Int{i}
}

func NewIntFromString(s string) Int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return Int{i}
}

func (i *Int) Int() int {
	return i.int
}

func (i *Int) ToString() string {
	return strconv.Itoa(i.int)
}

func (i *Int) Increment() {
	i.int++
}

func (i *Int) IncrementBy(by int) {
	i.int += by
}

func (i *Int) Decrement() {
	i.int--
}

func (i *Int) DecrementBy(by int) {
	i.int -= by
}

func (i *Int) Equals(other Int) bool {
	return i.int == other.int
}

func (i *Int) EqualTo(other int) bool {
	return i.int == other
}

func (i *Int) IsLarger(other Int) bool {
	return i.int > other.int
}

func (i *Int) IsLargerThan(other int) bool {
	return i.int > other
}

func (i *Int) IsSmaller(other Int) bool {
	return i.int < other.int
}

func (i *Int) IsSmallerThan(other int) bool {
	return i.int < other
}

func (i *Int) MustBePositive() {
	if i.int < 0 {
		panic("Int must be positive")
	}
}

func (i *Int) MustBeNegative() {
	if i.int > 0 {
		panic("Int must be negative")
	}
}

func (i *Int) MustBeZero() {
	if i.int != 0 {
		panic("Int must be zero")
	}
}

func (i *Int) MustEqual(other Int) {
	if i.int != other.int {
		panic(fmt.Sprintf("Int must equal %d", other.int))
	}
}

func (i *Int) MustBeEqualTo(other int) {
	if i.int != other {
		panic(fmt.Sprintf("Int must equal %d", other))
	}
}
