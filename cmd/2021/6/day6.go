package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
	"strings"
)

func main() {
	input := util.ReadFile("inputs/2021/day6.txt")
	part1(input, 80)
	part2(input, 256)
}

func part1(input string, days int) {
	fishCountersRaw := strings.Split(input, ",")
	fish := make([]int, len(fishCountersRaw))
	for i, s := range fishCountersRaw {
		fish[i] = util.ParseInt(s)
	}

	fmt.Println("Part1: Initial state:", fish)
	for i := 0; i < days; i++ {
		for j, day := range fish {
			if day == 0 {
				// reset counter
				fish[j] = 6
				// add a new fish
				fish = append(fish, 8)
			} else {
				// elapse 1 day
				fish[j] = day - 1
			}
		}
	}
	fmt.Printf("Part1: After %d days there are %d lanternfish!\n", days, len(fish))
}

func part2(input string, days int) {
	// part 2 isn't possible using the same naive method
	// instead of simulating fish one at a time, let's try simulating them in buckets
	// depending on the number of days left to spawn

	// holds the count of young and old fish, the age of the fish distinguished by the location in the slice
	youngFishCount := make([]int64, 9)
	oldFishCount := make([]int64, 7)

	fishCountersRaw := strings.Split(input, ",")
	for _, s := range fishCountersRaw {
		oldFishCount[util.ParseInt(s)]++
	}
	total := func() int64 {
		total := int64(0)
		for _, v := range youngFishCount {
			total += v
		}
		for _, v := range oldFishCount {
			total += v
		}
		return total
	}

	// offsets represent the 0 day in each count slice
	youngoffset, oldOffset := 0, 0
	for i := 0; i < days; i++ {
		// young fish have their first spawn
		youngSpawn := youngFishCount[youngoffset]
		// old fish have another spawn
		oldSpawn := oldFishCount[oldOffset]

		// young fish are now old fish, the spawns keep the number of young fish the same
		oldFishCount[oldOffset] += youngSpawn
		// old fish create new young fish
		youngFishCount[youngoffset] += oldSpawn

		// this is equivalent to aging the fish by 1 day
		youngoffset = (youngoffset + 1) % 9
		oldOffset = (oldOffset + 1) % 7

		fmt.Printf("Part2: After %d days there are %d lanternfish!\n", i+1, total())
		fmt.Printf("[Debug] yOffset=%d; oOffset=%d; young=%d; old=%d\n", youngoffset, oldOffset, youngFishCount, oldFishCount)
	}
}
