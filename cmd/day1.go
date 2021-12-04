package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := input()
	a1 := part1(input)
	fmt.Println("A1:", a1)
	a2 := part2(input)
	fmt.Println("A2:", a2)
}

func input() []int {
	dat, err := os.ReadFile("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	vals := []int{}
	for _, vs := range strings.Split(string(dat), "\r\n") {
		v, err := strconv.Atoi(vs)
		if err != nil {
			panic(err)
		}
		vals = append(vals, v)
	}
	return vals
}

func part1(input []int) int {
	count := 0
	prev := -1
	for _, v := range input {
		if v > prev {
			count++
		}
		prev = v
	}
	return count - 1
}

func part2(input []int) int {
	// the measurements in common between 2 adjacent sums can be ignored
	// and simplified by comparing the first measurement in sum A with the last
	// measurement in sum B, etc.
	count := 0
	inLen := len(input)
	for i := 3; i < inLen; i++ {
		v1 := input[i-3]
		v2 := input[i]
		if v2 > v1 {
			count++
		}
	}
	return count
}
