package cool

func Try(f func() AnyArray) (AnyArray, Any) {
	var err Any = nil
	var result AnyArray = nil

	func() {
		defer func() {
			err = recover()
		}()

		result = f()
	}()

	return result, err
}
