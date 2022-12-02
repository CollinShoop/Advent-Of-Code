package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
	"strings"
)

func main() {
	manual := NewPaperManual(util.ReadLines("inputs/2021/day13.txt"))
	manual.fold()
	fmt.Printf("(Part1) There are %d dots visible after the first fold\n", len(manual.dots))

	// Example manual print for part 2: LRFJBJEH

	// #....###..####...##.###....##.####.#..#.
	// #....#..#.#.......#.#..#....#.#....#..#.
	// #....#..#.###.....#.###.....#.###..####.
	// #....###..#.......#.#..#....#.#....#..#.
	// #....#.#..#....#..#.#..#.#..#.#....#..#.
	// ####.#..#.#.....##..###...##..####.#..#.

	for manual.fold() {
	}
	fmt.Println("(Part2)")
	manual.print()
}

type foldingManual struct {
	dots  map[vector2]struct{}
	size  vector2
	folds []fold
}

type vector2 struct {
	x, y int
}

type fold struct {
	axis string
	v    int
}

func NewPaperManual(lines []string) *foldingManual {
	p := &foldingManual{
		dots:  map[vector2]struct{}{},
		folds: []fold{},
		size:  vector2{x: 0, y: 0},
	}
	for _, line := range lines {
		// parse dot coordinate
		if strings.Contains(line, ",") {
			dotTerms := strings.Split(line, ",")
			v2 := vector2{
				x: util.ParseInt(dotTerms[0]),
				y: util.ParseInt(dotTerms[1]),
			}
			// keep track of the size of the board according to the maximum of each axis
			// this is just for getting a nice print at the end without any guess work
			if v2.x > p.size.x {
				p.size.x = v2.x
			}
			if v2.y > p.size.y {
				p.size.y = v2.y
			}
			p.dots[v2] = struct{}{}
		}

		// parse flip instruction
		if strings.Contains(line, "fold along") {
			foldTerms := strings.Split(strings.TrimPrefix(line, "fold along "), "=")
			p.folds = append(p.folds, fold{
				axis: foldTerms[0],
				v:    util.ParseInt(foldTerms[1]),
			})
		}
	}
	return p
}

func (fm *foldingManual) fold() bool {
	if len(fm.folds) == 0 {
		return false
	}
	// pop the next fold instruction
	f := fm.folds[0]
	fm.folds = fm.folds[1:]

	// flips a value v over a fold value f
	flipV := func(v int, f int) int {
		if v < f {
			return v
		}
		return f - (v - f)
	}
	// flips a vector2 v over by a fold f
	flipV2 := func(v vector2, f fold) vector2 {
		if f.axis == "y" {
			return vector2{
				x: v.x,
				y: flipV(v.y, f.v),
			}
		}
		return vector2{
			x: flipV(v.x, f.v),
			y: v.y,
		}
	}
	// flip each point into a new map
	folded := map[vector2]struct{}{}
	for d := range fm.dots {
		folded[flipV2(d, f)] = struct{}{}
	}
	fm.dots = folded

	// update size accordingly
	// this is just to get a nice print-out at the end
	if f.axis == "x" {
		fm.size.x = fm.size.x / 2
	} else if f.axis == "y" {
		fm.size.y = fm.size.y / 2
	}
	return true
}

func (fm *foldingManual) print() {
	for y := 0; y < fm.size.y; y++ {
		for x := 0; x < fm.size.x; x++ {
			_, ok := fm.dots[vector2{x: x, y: y}]
			if ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
