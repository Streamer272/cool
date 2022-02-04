package main

import (
	"fmt"
	"github.com/Streamer272/cool/async"
	"github.com/Streamer272/cool/types"
	"time"
)

func longRunningTask() types.AnyArray {
	time.Sleep(time.Second * 2)
	return types.AnyArray{5}
}

func failingTask() types.AnyArray {
	panic("Failed")
	return types.AnyArray{5}
}

func main() {
	f1 := async.Async(longRunningTask)

	fmt.Println("awaited", async.Await(f1))
	f1.Then(func(result types.AnyArray) {
		fmt.Println("then", result)
	})

	f2 := async.Async(longRunningTask)

	fmt.Println("awaited", async.Await(f2))
	f2.Catch(func(err types.Error, result types.AnyArray) {
		fmt.Println("catch", err, "result", result)
	})

	f3 := async.Async(failingTask)
	f3.Catch(func(err types.Error, result types.AnyArray) {
		fmt.Println("catch", err, "result", result)
	})

	time.Sleep(time.Second * 5)
}
