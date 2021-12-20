package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
)

func main() {
	/*
			This solution likely isn't optimal, but it's not too bad. At least it's going to be much quicker
			than a recursive depth-first approach without any dynamic-programming.

			The basic idea is to use an iterative dynamic approach to calculate the total risk
		    of travel to each point on the map, and then settling the risk values by iteratively re-calculating
		    the risk to each cell according to neighboring cells, until no changes are made.
			Once no changes are made, we know the map is complete, so the risk of the bottom-right cell has to be accurate.
	*/

	cm := NewChitonMap(util.ReadLines("inputs/day15.txt"))
	for cm.Evolve() > 0 {
		//cm.Print()
	}
	fmt.Println("(Part1) The safest risk is", cm.End().riskTotal)

	fmt.Println("")

	cm2 := NewExpandedChitonMap(cm)
	for cm2.Evolve() > 0 {
	}
	fmt.Println("(Part2) The safest risk is", cm2.End().riskTotal)
}

func NewChitonMap(inputs []string) *chitonMap {
	cells := make([][]*cell, len(inputs))
	for i, row := range inputs {
		cells[i] = make([]*cell, len(row))
		for j := range row {
			cells[i][j] = &cell{
				value:     int(row[j] - '0'),
				riskTotal: -1, // unknown
			}
		}
	}
	cells[0][0].riskTotal = 0 // 0,0 is the only cell with known risk: 0
	return &chitonMap{
		cells: cells,
	}
}

func NewExpandedChitonMap(cm *chitonMap) *chitonMap {
	height := len(cm.cells)
	width := len(cm.cells[0])

	cells := make([][]*cell, height*5)
	for i := 0; i < len(cm.cells)*5; i++ {
		cells[i] = make([]*cell, width*5)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			originalCell := cm.cells[y][x]
			// project out the original cell to each of the 5x5 expanded maps
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					cells[y+i*height][x+j*width] = &cell{
						// there's some extra maths with -1, %9, +1 here to wrap around 10 to 1, etc.
						value:     (originalCell.value+i+j-1)%9 + 1,
						riskTotal: -1,
					}
				}
			}
		}
	}

	cells[0][0].riskTotal = 0 // 0,0 is the only cell with known risk: 0
	return &chitonMap{
		cells: cells,
	}
}

func (cm *chitonMap) Evolve() int {
	updateCount := 0
	for i, row := range cm.cells {
		for j := range row {
			updateCount += cm.update(i, j)
		}
	}
	fmt.Println("Evolution resulted in", updateCount, "updates with an end risk of", cm.End().riskTotal)
	return updateCount
}

func (cm *chitonMap) Print() {
	for _, row := range cm.cells {
		for _, c := range row {
			fmt.Printf("%d", c.value)
		}
		fmt.Printf("  |   ")
		for _, c := range row {
			fmt.Printf("%04d ", c.riskTotal)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (cm *chitonMap) End() *cell {
	return cm.cells[len(cm.cells)-1][len(cm.cells[0])-1]
}

func (cm *chitonMap) update(x, y int) int {
	updateCount := 0
	inBounds := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < len(cm.cells) && y < len(cm.cells[0])
	}
	updateIfSafer := func(x1, y1, x2, y2 int) {
		if !inBounds(x1, y1) || !inBounds(x2, y2) {
			return
		}
		c1 := cm.cells[x1][y1]
		c2 := cm.cells[x2][y2]
		if c1.riskTotal == -1 {
			return // can't evaluate riskTotal to a cell that has unknown riskTotal
		}
		if c1.riskTotal+c2.value < c2.riskTotal || c2.riskTotal == -1 {
			updateCount++
			c2.riskTotal = c1.riskTotal + c2.value
		}
	}
	updateIfSafer(x, y, x+1, y)
	updateIfSafer(x, y, x, y+1)
	updateIfSafer(x, y, x-1, y)
	updateIfSafer(x, y, x, y-1)
	return updateCount
}

type chitonMap struct {
	cells [][]*cell
}

type cell struct {
	value     int
	riskTotal int
}
