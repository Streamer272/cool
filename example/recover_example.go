package main

import "github.com/Streamer272/cool/recover"

func main() {
	defer recover.Recover()

	panic("test err")
}
