package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
)

func main() {
	input := util.ReadInts("inputs/2021/day1.txt")
	a1 := part1(input)
	fmt.Println("A1:", a1)
	a2 := part2(input)
	fmt.Println("A2:", a2)
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
