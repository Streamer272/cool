package builtin

import (
	"fmt"
	"github.com/Streamer272/cool/types"
	"time"
)

func TimeDecorator(f func()) func() types.AnyArray {
	return func() types.AnyArray {
		// time function duration
		start := time.Now()
		f()
		end := time.Since(start)

		fmt.Printf("Function took %s\n", end.String())

		return types.AnyArray{end}
	}
}
