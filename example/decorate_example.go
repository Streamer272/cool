package main

import (
	"fmt"
	"github.com/Streamer272/cool/decorate"
	"github.com/Streamer272/cool/decorate/builtin"
	"github.com/Streamer272/cool/types"
	"time"
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

var testTime = decorate.Decorate(func() {
	fmt.Println("test time")
	time.Sleep(time.Millisecond * 200)
}, builtin.TimeDecorator)

func main() {
	testPrint()
	testError()
	testTime()
}
