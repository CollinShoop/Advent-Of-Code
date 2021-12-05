package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
	"strings"
)

func main() {
	inputs := util.ReadLines("inputs/day4.txt")
	moves := strings.Split(inputs[0], ",")
	boards := parseBoards(inputs[1:])

	bingoCount := 0
	for moveN, move := range moves {
		for boardN, board := range boards {
			if bingo, _ := board.bingo(); bingo {
				// skip boards that already have bingo
				continue
			}
			board.move(move)
			if bingo, score := board.bingo(); bingo {
				bingoCount++
				if bingoCount == 1 {
					// First Bingo
					board.print()
					fmt.Printf("First Bingo on move #%d (%s) is board %d with a score of %d\n", moveN, move, boardN, score*util.ParseInt(move))
				} else if bingoCount == len(boards) {
					// Last Bingo
					board.print()
					fmt.Printf("Last Bingo on move #%d (%s) is board %d with a score of %d\n", moveN, move, boardN, score*util.ParseInt(move))
				}
			}
		}
	}

}

type board struct {
	moves int
	tiles [][]string
}

func (b *board) print() {
	for _, row := range b.tiles {
		fmt.Println(row)
	}
}

func (b *board) move(n string) {
	b.moves++
	for i, row := range b.tiles {
		for j, v := range row {
			if v == n {
				b.tiles[i][j] = "#"
			}
		}
	}
}

func (b *board) bingo() (bingo bool, score int) {
	for i := 0; i < 5; i++ {
		row, col := true, true
		for j := 0; j < 5; j++ {
			if b.tiles[i][j] != "#" {
				score += util.ParseInt(b.tiles[i][j])
				row = false
			}
			if b.tiles[j][i] != "#" {
				col = false
			}
		}
		if row || col {
			bingo = true
		}
	}
	return
}

func parseBoards(inputs []string) []*board {
	boards := []*board{}
	for i := 0; i < len(inputs); i += 6 {
		puzzleInput := inputs[i+1 : i+6]
		b := &board{
			tiles: [][]string{},
		}
		for _, line := range puzzleInput {
			row := []string{}
			for j := 0; j < 14; j += 3 {
				row = append(row, strings.TrimSpace(line[j:j+2]))
			}
			b.tiles = append(b.tiles, row)
		}
		boards = append(boards, b)
	}
	return boards
}
