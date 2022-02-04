package async

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

func (thread *Thread) Then(callback func(array types.AnyArray)) {
	go func() {
		callback(thread.callable())
	}()
}

func (thread *Thread) Catch(callback func(types.Error, types.AnyArray)) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				callback(err, types.AnyArray{err})
			}
		}()

		callback(nil, thread.callable())
	}()
}
