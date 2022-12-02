package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
	"sort"
	"strings"
	"time"
)

func main() {
	part1()

	start := time.Now()
	part2()
	fmt.Println("(Part2) Took", time.Since(start))
}

func part1() {
	sum := 0
	for _, input := range util.ReadLines("inputs/2021/day8.txt") {
		for _, digit := range strings.Split(strings.Split(input, " | ")[1], " ") {
			switch len(digit) {
			//   1, 4, 7, 8
			case 2, 4, 3, 7:
				sum++
			}
		}
	}
	fmt.Println("(Part1) Sum is", sum)
}

func part2() {
	// sorts a string by its runes
	// example: runeSort("cat") -> "act"
	runeSort := func(s string) string {
		r := strings.Split(s, "")
		sort.Strings(r)
		return strings.Join(r, "")
	}
	// produces a new slice of ss minus
	// example: minus(["a", "b", "c"], "b") -> ["a", "c"]
	minus := func(ss []string, ex string) []string {
		r := make([]string, len(ss)-1)[:0]
		for _, s := range ss {
			if s != ex {
				r = append(r, s)
			}
		}
		return r
	}
	// finds all the inputs matching a desired length
	// example: matchByLen(["aaa", "bb", "ccc", "d"], 3) -> ["aaa", "ccc"]
	matchByLen := func(s []string, desiredLen int) []string {
		matches := []string{}
		for _, r := range s {
			if len(r) == desiredLen {
				matches = append(matches, r)
			}
		}
		return matches
	}
	// counts the number of runes in common between strA and strB, assuming no repeating runes.
	// example: countRunesInCommon("tac", "cats") -> 3
	countRunesInCommon := func(strA, strB string) int {
		commonCount := 0
		counts := make([]int, 7)
		for _, r := range strA {
			counts[r-'a']++
		}
		for _, r := range strB {
			if counts[r-'a'] > 0 {
				commonCount++
			}
		}
		return commonCount
	}
	// finds the first instance of ss with the desired number of runes in common with the target
	// example: matchByCommonRuneCount(["rant", "sing"], "smart", 1) -> "sing"
	matchByCommonRuneCount := func(ss []string, target string, count int) string {
		for _, s := range ss {
			if countRunesInCommon(s, target) == count {
				return s
			}
		}
		panic(fmt.Sprintf("cant find match of count %d comparing %s to %s", count, target, ss))
	}

	sum := 0

	inputs := util.ReadLines("inputs/2021/day8.txt")
	for _, input := range inputs {
		fmt.Printf("(Part2) Inputs: '%s'\n", input)
		parts := strings.Split(input, " | ")
		readings := strings.Split(parts[0], " ")
		for i, s := range readings {
			readings[i] = runeSort(s)
		}

		// slice of known digit mappings where the index is the digit (ex [1] = "cf")
		displayMapping := make([]string, 10)

		// digits 1, 4, 7, and 8 are uniquely identified by length
		displayMapping[1] = matchByLen(readings, 2)[0]
		displayMapping[4] = matchByLen(readings, 4)[0]
		displayMapping[7] = matchByLen(readings, 3)[0]
		displayMapping[8] = matchByLen(readings, 7)[0]

		// The remaining digits can be found by process of elimination, using a
		// static algorithm, by looking for specific overlap between readings
		// reading and known displays of the same length.

		// [Find 3] Take readings of size 5 (2, 3, 5), compare with 1. The reading with 2 in common is 3.
		readings5 := matchByLen(readings, 5)
		displayMapping[3] = matchByCommonRuneCount(readings5, displayMapping[1], 2)
		// [Find 2] Take readings of size 5 (2, 3, 5), compare with 4. The reading with 2 in common is 2.
		displayMapping[2] = matchByCommonRuneCount(readings5, displayMapping[4], 2)
		// [Find 5] Take readings of size 5 minus 2 (3, 5), compare with 1. The reading with 1 in common is 5.
		readings5 = minus(readings5, displayMapping[2])
		displayMapping[5] = matchByCommonRuneCount(readings5, displayMapping[1], 1)
		// [Find 9] Take readings of size 6 (0, 6, 9), compare with 3. The reading with 5 in common is 9.
		readings6 := matchByLen(readings, 6)
		displayMapping[9] = matchByCommonRuneCount(readings6, displayMapping[3], 5)
		// [Find 0] Take readings of size 6 minus 9 (0, 6), compare with 1. The reading with 2 in common is 0.
		readings6 = minus(readings6, displayMapping[9])
		displayMapping[0] = matchByCommonRuneCount(readings6, displayMapping[1], 2)
		// [Find 6] As the last remaining.
		readings6 = minus(readings6, displayMapping[0])
		displayMapping[6] = readings6[0]

		// reverse the display mapping lookup direction
		lookupDigitByDisplay := map[string]int{}
		for i, display := range displayMapping {
			lookupDigitByDisplay[display] = i
		}

		output := strings.Split(parts[1], " ")
		outputValue := 0
		for _, s := range output {
			digit, ok := lookupDigitByDisplay[runeSort(s)]
			if !ok {
				panic(fmt.Sprintf("can't find output '%s' in lookup %v", runeSort(s), lookupDigitByDisplay))
			}
			fmt.Printf("\tMapped output display %s (%s) to digit %d\n", s, runeSort(s), digit)
			// append numerically
			outputValue = (outputValue * 10) + digit
		}
		fmt.Printf("\tMapped outputs %s to value %d\n", output, outputValue)
		sum += outputValue
	}
	fmt.Println("(Part2) Sum is", sum)
}
