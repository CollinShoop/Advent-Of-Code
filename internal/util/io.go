package util

import (
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) string {
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func ReadLines(path string) []string {
	return strings.Split(ReadFile(path), "\r\n")
}

func ReadInts(path string) []int {
	vals := []int{}
	for _, vs := range ReadLines(path) {
		v, err := strconv.Atoi(vs)
		if err != nil {
			panic(err)
		}
		vals = append(vals, v)
	}
	return vals
}
