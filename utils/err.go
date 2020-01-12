package utils

func CheckNoError(err error) {
	if err != nil {
		panic(err)
	}
}
func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}
