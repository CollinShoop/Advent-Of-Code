package main

import (
	"errors"
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
	"strconv"
	"strings"
)

func main() {
	inputs := util.ReadLines("inputs/day2.txt")
	part1(inputs)
	part2(inputs)
}

func part1(inputs []string) {
	hpos, depth := 0, 0
	for _, line := range inputs {
		direction, v := parseLine(line)
		switch direction {
		case "forward":
			hpos += v
		case "up":
			depth -= v
		case "down":
			depth += v
		default:
			panic(errors.New("unexpected input direction: " + line))
		}
	}
	fmt.Printf("Part 1 hpos=%d; depth=%d; answer=%d\n", hpos, depth, hpos*depth)
}

func part2(inputs []string) {
	hpos, depth, aim := 0, 0, 0
	for _, line := range inputs {
		direction, v := parseLine(line)
		switch direction {
		case "forward":
			hpos += v
			depth += aim * v
		case "up":
			aim -= v
		case "down":
			aim += v
		default:
			panic(errors.New("unexpected input direction: " + line))
		}
	}
	fmt.Printf("Part 2 hpos=%d; depth=%d; answer=%d\n", hpos, depth, hpos*depth)
}

func parseLine(line string) (string, int) {
	tokens := strings.Split(line, " ")
	if len(tokens) != 2 {
		panic(errors.New("Unexpected line: " + line))
	}
	v, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}
	return tokens[0], v
}
