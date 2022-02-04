package main

import (
	"github.com/Streamer272/cool/exit"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		exit.ExitFailure()
	} else {
		exit.ExitSuccess()
	}
}
