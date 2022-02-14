package check

import (
	"os"
	"fmt"
)

func Check(err interface{}) {
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		os.Exit(1)
	}
}
