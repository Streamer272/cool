package exit

import (
	"fmt"
	"os"
)

func Exit(code int) {
	fmt.Printf("Exiting with code %v\n", code)
	os.Exit(code)
}

func ExitSuccess() {
	Exit(0)
}

func ExitFailure() {
	Exit(1)
}
