package decorate

import "github.com/Streamer272/cool/types"

func Decorate(f func(), dec func(func()) func() types.AnyArray) func() types.AnyArray {
	return dec(f)
}
