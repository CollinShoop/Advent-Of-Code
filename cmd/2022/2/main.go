package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
)

// https://adventofcode.com/2022/day/2
func main() {
	input := util.ReadLines("inputs/2022/day2.txt")
	fmt.Println("A1:", part1(input))
	fmt.Println("A2:", part2(input))
}

// What would your total score be if everything goes exactly according to your strategy guide?
func part1(input []string) int {
	totalScore := 0
	for _, s := range input {
		if s == "" {
			continue
		}
		opponentMove := int(s[0]) - 'A'
		myMove := int(s[2]) - 'X'
		totalScore += score(myMove, opponentMove)
	}
	return totalScore
}

// Following the Elf's instructions for the second column, what would your total score be if everything
// goes exactly according to your strategy guide?
func part2(input []string) int {
	totalScore := 0
	for _, s := range input {
		if s == "" {
			continue
		}
		opponentMove := int(s[0]) - 'A'
		outcome := int(s[2]) - 'X' // 0=lose, 1=draw, 2=win
		myMove := predict(opponentMove)[outcome]
		totalScore += score(myMove, opponentMove)
	}
	return totalScore
}

// Given a move (0 = Rock, 1 = Paper, 2 = Scissors), returns a list of moves that would result int
// [lose, draw, win]

// Moves are of the form:
// 0 = rock, 1 = paper, 2 = scissors
func predict(move int) []int {
	return []int{
		(move + 2) % 3,
		move,
		(move + 1) % 3,
	}
}

// Given two rock-paper-scissors moves, returns the score
// of the outcome for player A.
//
// Moves are of the form:
// 0 = rock (1 point), 1 = paper (2 points), 2 = scissors (3 points)
//
// Outcomes:
// Lose = 0 points, Draw = 3 pints, Win = 6 points
func score(a int, b int) int {
	if (a+1)%3 == b {
		return a + 1 // b wins
	} else if (b+1)%3 == a {
		return a + 7 // a wins
	}
	return a + 4 // draw
}
