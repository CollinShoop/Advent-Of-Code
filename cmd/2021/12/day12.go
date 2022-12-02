package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
	"strings"
)

func main() {
	cave := NewCave(util.ReadLines("inputs/2021/day12.txt"))

	cave.Explore(false)
	fmt.Printf("(Part1) There are %d paths\n", cave.PathCount)

	cave.Explore(true)
	fmt.Printf("(Part2) With a sinle twice visit, there are %d paths\n", cave.PathCount)
}

type Cave struct {
	caveMap         map[string][]string
	visited         map[string]int
	allowTwiceVisit bool
	twiceVisited    bool

	PathCount int
}

func NewCave(inputs []string) *Cave {
	c := &Cave{
		caveMap: map[string][]string{},
	}
	connect := func(from, to string) {
		paths, ok := c.caveMap[from]
		if ok {
			c.caveMap[from] = append(paths, to) // add connection from known cave to another
		} else {
			c.caveMap[from] = []string{to} // add connection from new cave to another
		}
	}
	for _, input := range inputs {
		caves := strings.Split(input, "-")
		connect(caves[0], caves[1])
		connect(caves[1], caves[0])
	}
	return c
}

// Explore starts the cave exploration, with the goal of finding the total number of unique paths
// through the cave system while following the cave-size rules. allowTwiceVisit specifies whether or not
// a single small cave may be visited twice.
//
// Exploration is done via depth-first backtracking by keeping track of which caves have and have not yet been visited
// the allowed number of times.
func (c *Cave) Explore(allowTwiceVisit bool) {
	c.allowTwiceVisit = allowTwiceVisit
	c.PathCount = 0
	c.visited = map[string]int{
		"start": 2,
	}
	for _, to := range c.caveMap["start"] {
		c.traverse(to)
	}
}

func (c *Cave) traverse(from string) {
	if !c.visit(from) {
		return
	}
	if from == "end" {
		c.PathCount++
	} else {
		for _, to := range c.caveMap[from] {
			c.traverse(to)
		}
	}
	c.backVisit(from)
}

func (c *Cave) visit(cave string) bool {
	if cave == strings.ToUpper(cave) { // lazy isBig
		return true
	}
	t := c.visited[cave]
	if t == 0 {
		c.visited[cave] = 1
		return true
	}
	// secret sauce for part 2, allow small cave to be visited twice
	// iff one already hasn't been
	if t == 1 && c.allowTwiceVisit && !c.twiceVisited {
		c.visited[cave] = 2
		c.twiceVisited = true
		return true
	}
	return false
}

func (c *Cave) backVisit(cave string) {
	t := c.visited[cave]
	c.visited[cave]--
	if t == 2 {
		c.twiceVisited = false
	}
}
