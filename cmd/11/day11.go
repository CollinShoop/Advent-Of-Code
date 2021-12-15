package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
)

var FLASHED = -1

func main() {
	inputs := util.ReadLines("inputs/day11.txt")

	// parse inputs into 2d grid
	grid := make([][]int, len(inputs))
	for i, row := range inputs {
		grid[i] = make([]int, len(row))
		for j := range row {
			grid[i][j] = int(row[j] - '0')
		}
	}

	// keep track of the number of steps elapsed and total number of Octopus flashes
	steps := 0
	flashCount := 0

	// flash and increment work together, recursively, to control the propagation
	// of energy levels within the grid. Once an Octopus flashes, the cell is marked with FLASHED (-1)
	// to indicate that the cell can't be energized further within the same turn (or in other words, the same
	// recursion tree).
	var flash func(i, j int)
	increment := func(i, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return
		}
		if grid[i][j] == FLASHED {
			return
		}
		grid[i][j]++
		if grid[i][j] > 9 {
			flash(i, j)
		}
	}
	flash = func(i, j int) {
		flashCount++
		grid[i][j] = FLASHED
		// above
		increment(i-1, j-1)
		increment(i-1, j)
		increment(i-1, j+1)
		// sides
		increment(i, j-1)
		increment(i, j+1)
		// below
		increment(i+1, j-1)
		increment(i+1, j)
		increment(i+1, j+1)
	}
	runStep := func() {
		steps++
		// gain 1 energy per turn
		for i := range grid {
			for j := range grid[i] {
				increment(i, j)
			}
		}
		// reset energy for any that flashed
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] == FLASHED {
					grid[i][j] = 0
				}
			}
		}
	}

	for i := 0; i < 100; i++ {
		runStep()
	}
	fmt.Printf("(Part1) After 100 steps, total of %d flashes\n", flashCount)

	for i := 0; i < 5000; i++ {
		lastFlashCount := flashCount
		runStep()
		// Detect synchronized flash - when the diff in flash count is equal to the number of Octopus
		if flashCount-lastFlashCount == len(inputs)*len(inputs[0]) {
			fmt.Printf("(Part2) First synchronized flash took place after %d steps\n", steps)
			break
		}
	}
}
