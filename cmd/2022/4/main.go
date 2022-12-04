package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
	"strings"
)

// https://adventofcode.com/2022/day/3
func main() {
	input := util.ReadLines("inputs/2022/day4.txt")
	fmt.Println("A1:", part1(input))
	fmt.Println("A2:", part2(input))
}

func part1(input []string) int {
	count := 0
	for _, line := range input {
		a1, a2, b1, b2 := parseRanges(line)
		// one range contained in the other
		if (a1 >= b1 && a2 <= b2) || (b1 >= a1 && b2 <= a2) {
			count++
		}
	}
	return count
}

func parseRanges(s string) (a1, a2, b1, b2 int) {
	ranges := strings.Split(s, ",")
	parseRange := func(s string) (a1, a2 int) {
		nums := strings.Split(s, "-")
		return util.ParseInt(nums[0]), util.ParseInt(nums[1])
	}
	a1, a2 = parseRange(ranges[0])
	b1, b2 = parseRange(ranges[1])
	return
}

func part2(input []string) int {
	count := 0
	for _, line := range input {
		a1, a2, b1, b2 := parseRanges(line)
		// one range contained in the other
		if (a1 >= b1 && a2 <= b2) || (b1 >= a1 && b2 <= a2) ||
			// left or right partial overlap
			(a1 <= b1 && b1 <= a2) || (a1 <= b2 && b2 <= a2) {
			count++
		}
	}
	return count
}
