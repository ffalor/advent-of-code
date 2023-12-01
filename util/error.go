package util

// ErrPanic throws a panic if the err is not nil
func ErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
