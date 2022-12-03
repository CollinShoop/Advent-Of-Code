package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
	"unicode"
)

// https://adventofcode.com/2022/day/3
func main() {
	input := util.ReadLines("inputs/2022/day3.txt")
	fmt.Println("A1:", part1(input))
	fmt.Println("A2:", part2(input))
}

func part1(input []string) int {
	sum := 0
	for _, rucksack := range input {
		size := len(rucksack)
		sum += score(commonLetter2(rucksack[:size/2], rucksack[size/2:]))
	}
	return sum
}

func commonLetter2(a, b string) rune {
	counts := map[rune]int{}
	for _, r := range a {
		counts[r]++
	}
	for _, r := range b {
		if _, exists := counts[r]; exists {
			return r
		}
	}
	return 0
}

func score(r rune) int {
	if unicode.IsLower(r) {
		return int(r-'a') + 1
	}
	return int(r-'A') + 27
}

func part2(input []string) int {
	sum := 0
	for i := 0; i < len(input); i += 3 {
		sum += score(commonLetter3(input[i], input[i+1], input[i+2]))
	}
	return sum
}

func commonLetter3(a, b, c string) rune {
	counts := map[rune]int{}
	countOnce := func(s string) {
		localCounts := map[rune]int{}
		for _, r := range s {
			localCounts[r]++
		}
		for r := range localCounts {
			counts[r]++
		}
	}
	countOnce(a)
	countOnce(b)
	countOnce(c)

	// find the letter that occurs in all 3
	for r, count := range counts {
		if count == 3 {
			return r
		}
	}
	return 0
}
