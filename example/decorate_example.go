package main

import (
	"fmt"
	"github.com/Streamer272/cool/decorate"
	"github.com/Streamer272/cool/types"
)

var test = decorate.Decorate(func(...types.Any) { // this is base function
	fmt.Println("test function called")
}, func(f func(...types.Any)) func(...types.Any) { // this is a decorator
	return func(args ...types.Any) { // this is base function wrapper
		fmt.Println("running")
		f(args...)
		fmt.Println("done")
	}
})

func main() {
	test()
}
