package types

import "strconv"

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

func (i *Int) ToString() string {
	return strconv.Itoa(i.int)
}
