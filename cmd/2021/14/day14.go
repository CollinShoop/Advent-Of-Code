package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
	"math"
	"strings"
)

func main() {
	polymer := NewPolymer(util.ReadLines("inputs/2021/day14.txt"))

	for i := 0; i < 10; i++ {
		polymer.Grow()
	}
	fmt.Printf("(Part1) After step 10: score=%d\n", polymer.Score())

	for i := 0; i < 30; i++ {
		polymer.Grow()
	}
	fmt.Printf("(Part2) After step 40: score=%d\n", polymer.Score())
}

type polymer struct {
	// remember the first letter of the polymer template for scoring
	first string
	// polymer growth rules
	rules map[string]string
	// polymer template pair counts, eg "NN" -> 23
	counts map[string]int64
}

func NewPolymer(lines []string) *polymer {
	rules := map[string]string{}
	counts := map[string]int64{}

	template := lines[0]
	for i := 0; i < len(template)-1; i++ {
		counts[template[i:i+2]]++
	}

	for i := 2; i < len(lines); i++ {
		ruleParts := strings.Split(lines[i], " -> ")
		rules[ruleParts[0]] = ruleParts[1]
	}
	return &polymer{
		first:  template[0:1],
		rules:  rules,
		counts: counts,
	}
}

func (p *polymer) Grow() {
	// polymer growth is gone in aggregate, kind of like a dynamic-programming problem.
	// where each pair is grown according to the specified rules as a group instead of being simulated individually.
	// simulating them individually works fine for the first few steps, but quickly grows to a scale
	// that is no longer manageable.

	growCounts := map[string]int64{}
	for pair, count := range p.counts {
		if rule, ok := p.rules[pair]; ok {
			// if there's a rule for splitting this pair,
			// split into 2 new sets of pairs
			s1 := string(pair[0]) + rule
			s2 := rule + string(pair[1])
			growCounts[s1] += count
			growCounts[s2] += count
		} else {
			// carry the pair over
			growCounts[pair] = count
		}
	}
	p.counts = growCounts
}

func (p *polymer) Score() int64 {
	rcounts := map[string]int64{
		// the first polymer letter is the only one that won't be counted later,
		p.first: 1,
	}
	min := int64(math.MaxInt64)
	max := int64(0)
	// for each pair, explicitly add only the second letter counts
	// this is because the first letter of each pair is implicitly
	// overlapping as the second letter of another pair
	for pair, count := range p.counts {
		rcounts[string(pair[1])] += count
	}
	for _, count := range rcounts {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}
	return max - min
}
