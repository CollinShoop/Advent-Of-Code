package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inputs := util.ReadLines("inputs/2021/day5.txt")
	part1(inputs)
	part2(inputs)
	fmt.Println("Took", time.Since(start))
}

func part1(inputs []string) {
	overlapCount := 0

	// allocate 1000x1000 grid for keeping track of vents
	ventGrid := make([][]byte, 1000)
	for i := 0; i < 1000; i++ {
		ventGrid[i] = make([]byte, 1000)
	}
	markVent := func(x, y int) {
		ventGrid[x][y]++
		if ventGrid[x][y] == 2 {
			overlapCount++
		}
	}

	for _, input := range inputs {
		// parse input line as set of 2 coordinates
		x1, y1, x2, y2 := parseInput(input)
		// walk from point 1 to point 2, marking vents
		if x1 == x2 {
			if y1 < y2 {
				for y := y1; y <= y2; y++ {
					markVent(x1, y)
				}
			} else {
				for y := y2; y <= y1; y++ {
					markVent(x1, y)
				}
			}
		} else if y1 == y2 {
			if x1 < x2 {
				for x := x1; x <= x2; x++ {
					markVent(x, y1)
				}
			} else {
				for x := x2; x <= x1; x++ {
					markVent(x, y1)
				}
			}
		}
	}

	fmt.Println("(Part 1) Number of overlapping vents:", overlapCount)
}

func part2(inputs []string) {
	overlapCount := 0

	// allocate 1000x1000 grid for keeping track of vents
	ventGrid := make([][]byte, 1000)
	for i := 0; i < 1000; i++ {
		ventGrid[i] = make([]byte, 1000)
	}
	markVent := func(x, y int) {
		ventGrid[x][y]++
		if ventGrid[x][y] == 2 {
			overlapCount++
		}
	}

	for _, input := range inputs {
		// parse input line as set of 2 coordinates
		x1, y1, x2, y2 := parseInput(input)
		// walk from point 1 to point 2, marking vents
		dx, dy := normalize(x1, y1, x2, y2)
		x, y := x1, y1
		for {
			markVent(x, y)
			if x == x2 && y == y2 {
				break
			}
			x += dx
			y += dy
		}
	}

	fmt.Println("(Part 2) Number of overlapping vents:", overlapCount)
}

func normalize(x1, y1, x2, y2 int) (dx, dy int) {
	if x1 < x2 {
		dx = 1
	} else if x2 < x1 {
		dx = -1
	}
	if y1 < y2 {
		dy = 1
	} else if y2 < y1 {
		dy = -1
	}
	return
}

func parseInput(s string) (x1, y1, x2, y2 int) {
	// ex s="645,570 -> 517,570"
	parts := strings.Split(s, " -> ")
	parsePoint := func(s string) (x, y int) {
		parts := strings.Split(s, ",")
		return util.ParseInt(parts[0]), util.ParseInt(parts[1])
	}
	x1, y1 = parsePoint(parts[0])
	x2, y2 = parsePoint(parts[1])
	return
}
