package helper

func IfErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
