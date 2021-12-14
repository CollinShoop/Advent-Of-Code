package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
	"sort"
)

func main() {
	part1()
	part2()
}

var openMap = map[uint8]uint8{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}
var pointMap = map[uint8]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func part1() {
	inputs := util.ReadLines("inputs/day10.txt")

	rStack := []uint8{}
	push := func(r uint8) {
		rStack = append(rStack, r)
	}
	pop := func() uint8 {
		r := rStack[len(rStack)-1]
		rStack = rStack[:len(rStack)-1]
		return r
	}

	score := 0
	for _, input := range inputs {
		rStack = rStack[:0] // reset stack

		// iterate through the chunk string 1 rune at a time
		// - for each open rune, push to the stack
		// - for each close rune, check that the corresponding open rune is at the top of the stack
		//   - if something else was popped, the chunk is corrupted and scored
		for i := range input {
			r := input[i]
			rOpen, ok := openMap[r] // get opening rune for a closing rune
			if !ok {
				push(r) // r is an open rune, start a new chunk
			} else {
				// r is a close rune, attempt to close chunk
				if rOpen != pop() { // corrupted
					score += pointMap[r]
					break
				}
			}
		}
	}
	fmt.Println("(Part1) Score:", score)
}

func part2() {
	inputs := util.ReadLines("inputs/day10.txt")

	// simple rune stack
	stack := []uint8{}
	push := func(r uint8) {
		stack = append(stack, r)
	}
	pop := func() uint8 {
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return r
	}

	scores := []int{}
	for _, input := range inputs {
		stack = stack[:0] // reset stack
		isCorrupted := false

		// iterate through the chunk string 1 rune at a time
		// - for each open rune, push to the stack
		// - for each close rune, check that the corresponding open rune is at the top of the stack
		//   - if something else was popped, the chunk is corrupted and discarded
		for i := range input {
			r := input[i]
			rOpen, ok := openMap[r] // get opening rune for a closing rune
			if !ok {
				push(r) // r is an open rune, start a new chunk
			} else {
				// r is a close rune, attempt to close chunk
				if rOpen != pop() {
					isCorrupted = true
					break
				}
			}
		}
		if isCorrupted {
			continue // could also do this with a break label but not as readable imo
		}

		// score whatever is left on the stack
		// this is done by looking at the open chunks from right-to-left as opposed to scoring
		// inferred closing chunks from left-to-right.
		score := 0
		for len(stack) > 0 {
			score = (score * 5) + pointMap[pop()]
		}
		scores = append(scores, score)

	}
	sort.Ints(scores)
	fmt.Println("(Part2) Middle score:", scores[len(scores)/2])
}
