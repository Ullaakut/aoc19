package aocutils

// Check panics if the given error is not nil.
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
