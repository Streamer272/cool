package recover

import "fmt"

func Recover() {
	if err := recover(); err != nil {
		fmt.Printf("Recovered %v\n", err)
	}
}
