package cool

func Recover(callback func(err Any)) {
	if err := recover(); err != nil {
		callback(err)
	}
}
