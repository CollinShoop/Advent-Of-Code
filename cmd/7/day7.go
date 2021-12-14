package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
	"math"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	input := util.ReadFile("inputs/day7.txt")
	strPositions := strings.Split(input, ",")
	positions := make([]int, len(strPositions))

	max := 0
	for i, s := range strPositions {
		pos := util.ParseInt(s)
		positions[i] = pos
		if pos > max {
			max = pos
		}
	}

	fuelTotals := make([]int, max)
	for target := 0; target < max; target++ {
		for _, pos := range positions {
			if pos > target {
				fuelTotals[target] += pos - target
			} else {
				fuelTotals[target] += target - pos
			}
		}
	}

	// mind the min fuel
	minFuel := fuelTotals[0]
	for _, fuel := range fuelTotals {
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	fmt.Printf("(Part1) Min fuel is %d\n", minFuel)
}

func part2() {
	input := util.ReadFile("inputs/day7.txt")
	strPositions := strings.Split(input, ",")
	positions := make([]int, len(strPositions))

	max := 0
	for i, s := range strPositions {
		pos := util.ParseInt(s)
		positions[i] = pos
		if pos > max {
			max = pos
		}
	}

	fuelTotals := make([]int, max)
	for target := 0; target < max; target++ {
		for _, pos := range positions {
			var diff int
			if pos > target {
				diff = pos - target
			} else {
				diff = target - pos
			}

			// I forget what this formula is called, but it gives the cumulative sum of [1, diff]
			fuelTotals[target] += (int(math.Pow(float64(diff), 2)) + diff) / 2
		}
	}

	// mind the min fuel
	minFuel := fuelTotals[0]
	for _, fuel := range fuelTotals {
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	fmt.Printf("(Part2) Min fuel is %d\n", minFuel)
}
