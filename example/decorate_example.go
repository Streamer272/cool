package main

import (
	"fmt"
	"github.com/Streamer272/cool/decorate"
	"github.com/Streamer272/cool/decorate/builtin"
	"github.com/Streamer272/cool/types"
)

var testPrint = decorate.Decorate(func() { // this is base function
	fmt.Println("test function called")
}, func(f func()) func() types.AnyArray { // this is a decorator
	return func() types.AnyArray { // this is base function wrapper
		fmt.Println("running")
		f()
		fmt.Println("done")
		return types.AnyArray{}
	}
})

var testError = decorate.Decorate(func() {
	panic("test error")
}, builtin.RecoverDecorator)

func main() {
	testPrint()
	testError()
}
