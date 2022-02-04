package builtin

import (
	"github.com/Streamer272/cool/recover"
	"github.com/Streamer272/cool/types"
)

func RecoverDecorator(f func()) func() types.AnyArray {
	return func() types.AnyArray {
		defer recover.Recover()

		f()

		return types.AnyArray{}
	}
}
