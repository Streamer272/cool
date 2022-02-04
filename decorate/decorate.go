package decorate

import "github.com/Streamer272/cool/types"

func Decorate(f func(...types.Any), dec func(func(...types.Any)) func(...types.Any)) func(...types.Any) {
	return dec(f)
}
