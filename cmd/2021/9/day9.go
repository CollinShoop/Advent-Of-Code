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

func part1() {
	inputs := util.ReadLines("inputs/2021/day9.txt")

	isLower := func(r uint8, i, j int) bool {
		if i < 0 || j < 0 || i >= len(inputs) || j >= len(inputs[i]) {
			return true // [i][j] is out of bounds
		}
		return r < inputs[i][j]
	}
	isMin := func(r uint8, i, j int) bool {
		return isLower(r, i-1, j) && // up
			isLower(r, i+1, j) && // down
			isLower(r, i, j-1) && // left
			isLower(r, i, j+1) // right
	}
	sum := 0
	for i, row := range inputs {
		for j := range row {
			r := inputs[i][j]
			if isMin(r, i, j) {
				sum += int(r - '/') // convert height rune to height + 1
			}
		}
	}
	fmt.Printf("(Part1) Sum of all low points is: %d\n", sum)
}

func part2() {
	// convert input into a mutable map
	var heightMap [][]uint8
	{
		for _, row := range util.ReadLines("inputs/2021/day9.txt") {
			heightMap = append(heightMap, []uint8(row))
		}
	}
	isEdge := func(i, j int) bool {
		if i < 0 || j < 0 || i >= len(heightMap) || j >= len(heightMap[i]) {
			return true
		}
		return heightMap[i][j] == '9'
	}
	// recursive mutating search counts the size of a basin starting at any point
	// within a basin, counting each map location exactly once.
	var searchBasin func(i, j int) int
	searchBasin = func(i, j int) int {
		if isEdge(i, j) {
			return 0
		}
		heightMap[i][j] = '9'
		return 1 + searchBasin(i+1, j) + searchBasin(i-1, j) + searchBasin(i, j+1) + searchBasin(i, j-1)
	}

	// scan the map for basins
	var basins []int
	for i := range heightMap {
		for j := range heightMap[i] {
			if isEdge(i, j) { // avoids counting basins of size 0
				continue
			}
			basins = append(basins, searchBasin(i, j))
		}
	}

	// sort basin sizes descending for ease of access
	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})
	fmt.Printf("(Part2) Answer: %d\n", basins[0]*basins[1]*basins[2])
}
