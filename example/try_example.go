package main

import (
	"fmt"
	"github.com/Streamer272/cool/try"
	"github.com/Streamer272/cool/types"
)

func willFail() types.AnyArray {
	panic("fail")
	return types.AnyArray{"res1", "res2"}
}

func main() {
	result, err := try.Try(willFail)
	fmt.Printf("Got %v with error %v\n", result, err)
}
