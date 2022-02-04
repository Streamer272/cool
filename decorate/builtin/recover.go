package builtin

import (
	"fmt"
	"github.com/Streamer272/cool/recover"
	"github.com/Streamer272/cool/types"
)

func RecoverDecorator(f func()) func() types.AnyArray {
	return func() types.AnyArray {
		fmt.Printf("RecoverDecorator %p\n", f)
		defer recover.Recover()

		f()

		return types.AnyArray{}
	}
}
