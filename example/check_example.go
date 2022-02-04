package main

import (
	"errors"
	"github.com/Streamer272/cool/check"
)

func main() {
	check.Check(nil)                    // OK
	check.Check(errors.New("test err")) // ERROR
}
