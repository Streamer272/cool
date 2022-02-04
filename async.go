package main

import "github.com/Streamer272/cool/types"

type Thread struct {
	callable func() types.AnyArray
}

func Async(callable func() types.AnyArray) Thread {
	return Thread{callable: callable}
}

func Await(thread Thread) types.AnyArray {
	return thread.callable()
}

func (thread *Thread) Then(callback func(...types.Any)) {
	go func() {
		callback(thread.callable()...)
	}()
}

func (thread *Thread) Catch(callback func(bool, ...types.Any)) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				callback(false, err)
			}
		}()

		callback(true, thread.callable()...)
	}()
}
