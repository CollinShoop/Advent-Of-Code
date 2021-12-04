package util

func Panic(v interface{}) {
	if v != nil {
		panic(v)
	}
}
