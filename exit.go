package main

import "os"

func ExitSuccess() {
	os.Exit(0)
}

func ExitFailure() {
	os.Exit(1)
}

func Exit(code int) {
	os.Exit(code)
}
