package main

func Check(err interface{}) {
	if err != nil {
		panic(err)
	}
}
