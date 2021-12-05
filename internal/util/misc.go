package util

import "strconv"

func Panic(v interface{}) {
	if v != nil {
		panic(v)
	}
}

func ParseInt(s string) (v int) {
	v, err := strconv.Atoi(s)
	Panic(err)
	return
}
