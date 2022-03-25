package main

import (
	"fmt"
	"time"
)

func Decorate(f func(), dec func(func()) func() AnyArray) func() AnyArray {
	return dec(f)
}

func RecoverDecorator(f func(), callback func(err Any)) func() AnyArray {
	return func() AnyArray {
		defer Recover(callback)

		f()

		return AnyArray{}
	}
}

func TimeDecorator(f func()) func() AnyArray {
	return func() AnyArray {
		// time function duration
		start := time.Now()
		f()
		end := time.Since(start)

		fmt.Printf("Function took %s\n", end.String())

		return AnyArray{end}
	}
}
