package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
	"strconv"
)

func main() {
	inputs := util.ReadLines("inputs/2021/day3.txt")
	fmt.Println("Part 1")
	part1(inputs)
	fmt.Println("")
	fmt.Println("Part 2")
	part2(inputs)
}

func part1(inputs []string) {
	counts := make([]int, 12)
	for _, bs := range inputs {
		for i, bitc := range bs {
			if bitc == '1' {
				counts[i]++
			}
		}
	}
	fmt.Println("Counts", counts)

	// convert counts into gamma "binary" array
	for i, count := range counts {
		if count > len(inputs)/2 {
			counts[i] = 1
		} else {
			counts[i] = 0
		}
	}

	gammaB := countsToBinaryStr(counts, false)
	fmt.Println("Gamma (binary):", gammaB)
	gamma := parseBinary(gammaB)
	fmt.Println("Gamma:", gamma)

	epsilonB := countsToBinaryStr(counts, true)
	fmt.Println("Epsilon (binary):", epsilonB)
	epsilon := parseBinary(epsilonB)
	fmt.Println("Epsilon:", epsilon)

	fmt.Println("Power:", gamma*epsilon)
}

func part2(inputs []string) {
	reduceN := func(inputs []string, n int, leastCommon bool) []string {
		count1 := 0
		for _, bs := range inputs {
			if bs[n] == '1' {
				count1++
			}
		}
		count0 := len(inputs) - count1
		var mask byte // '0' or '1' to filter bit n
		if leastCommon {
			if count0 <= count1 {
				mask = '0'
			} else {
				mask = '1'
			}
		} else {
			if count1 >= count0 {
				mask = '1'
			} else {
				mask = '0'
			}
		}
		// find all inputs that match mask at bit n
		matches := []string{}
		for _, input := range inputs {
			if mask == input[n] {
				matches = append(matches, input)
			}
		}
		return matches
	}
	reduce := func(leastCommon bool) string {
		// reduce inputs 1 bit at a time, left to right, until only 1 match remains
		inputs := inputs
		for n := 0; n < 12; n++ {
			inputs = reduceN(inputs, n, leastCommon)
			if len(inputs) == 1 {
				return inputs[0]
			}
		}
		return ""
	}

	oxygenRating := reduce(false)
	c02Rating := reduce(true)
	fmt.Println("Oxygen Rating:", oxygenRating, parseBinary(oxygenRating))
	fmt.Println("C02 Rating:", c02Rating, parseBinary(c02Rating))
	fmt.Println("Life Support Rating:", parseBinary(oxygenRating)*parseBinary(c02Rating))
}

func countsToBinaryStr(input []int, invert bool) string {
	str := ""
	for _, b := range input {
		if invert {
			b = (b + 1) % 2
		}
		str += strconv.Itoa(b)
	}
	return str
}

func parseBinary(str string) int64 {
	v, err := strconv.ParseInt(str, 2, 64)
	util.Panic(err)
	return v
}
