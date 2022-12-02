package main

import (
	"container/heap"
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
)

// https://adventofcode.com/2022/day/1
func main() {
	input := util.ReadLines("inputs/2022/day1.txt")
	a1 := part1(input)
	fmt.Println("A1:", a1)
	a2 := part2(input)
	fmt.Println("A2:", a2)
}

// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
func part1(input []string) int {
	max := 0
	elfCal := 0
	for _, s := range input {
		// new elf
		if s == "" {
			if elfCal > max {
				max = elfCal
			}
			elfCal = 0
		} else {
			// count cals
			elfCal += util.ParseInt(s)
		}
	}
	return max
}

// How many Calories are those Elves carrying in total?
func part2(input []string) int {
	// should have just used a slice and sorted it
	top := util.NewIntHeap(false)
	heap.Init(top)

	elfCal := 0
	for _, s := range input {
		// new elf
		if s == "" {
			heap.Push(top, elfCal)
			elfCal = 0
		} else {
			// count cals
			elfCal += util.ParseInt(s)
		}
	}

	// grab top 3
	sum := 0
	for i := 0; i < 3; i++ {
		v := heap.Pop(top).(int)
		fmt.Println(v)
		sum += v
	}
	return sum
}
