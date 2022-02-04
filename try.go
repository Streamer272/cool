package main

import (
	"github.com/Streamer272/cool/types"
)

func Try(function func() types.AnyArray) (err types.Any, result types.AnyArray) {
	if err := recover(); err != nil {
		return err, types.AnyArray{}
	}

	return nil, function()
}
