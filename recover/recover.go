package recover

import "fmt"

func Recover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recovered %v\n", err)
		}
	}()
}
