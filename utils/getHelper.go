package utils

type GetHelper struct {
}

func MustGetString(v string, err error) string {
	if err != nil {
		panic(err)
	}
	return v
}

func MustGetInt(v int, err error) int {
	if err != nil {
		panic(err)
	}
	return v
}
