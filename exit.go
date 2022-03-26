package cool

import (
	"os"
)

func Exit(code int) {
	os.Exit(code)
}

func ExitSuccess() {
	Exit(0)
}

func ExitFailure() {
	Exit(1)
}
