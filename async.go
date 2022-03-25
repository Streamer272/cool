package main

type callable func() AnyArray

type Thread struct {
	callable callable
}

func Async(callable callable) Thread {
	return Thread{callable: callable}
}

func Await(thread Thread) AnyArray {
	return thread.callable()
}

func (thread *Thread) Then(callback func(array AnyArray)) {
	go func() {
		callback(thread.callable())
	}()
}

func (thread *Thread) Catch(callback func(Error, AnyArray)) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				callback(err, AnyArray{})
			}
		}()

		callback(nil, thread.callable())
	}()
}
