package try

import (
	"github.com/Streamer272/cool/types"
)

func Try(f func() types.AnyArray) (types.AnyArray, types.Any) {
	var err types.Any = nil
	var result types.AnyArray = nil

	func() {
		defer func() {
			err = recover()
		}()

		result = f()
	}()

	return result, err
}
