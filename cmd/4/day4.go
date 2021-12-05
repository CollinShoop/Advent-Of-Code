package main

import (
	"fmt"
	"github.com/collinshoop/adventofcode2021/internal/util"
	"strings"
	"time"
)

func main() {
	{
		fmt.Println("Basic Solution")
		start := time.Now()
		basicSolution()
		fmt.Println("Basic Solution Took", time.Since(start))
	}
	fmt.Println()
	{
		fmt.Println("Optimized Solution")
		start := time.Now()
		optimizedSolution()
		fmt.Println("Optimized Solution Took", time.Since(start))
	}
}

func optimizedSolution() {
	/*
		A summary of the approach taken here:
		1. re-map each bingo tile to the move in which that value is called
		2. find the number of moves needed to "Bingo!" each board by finding the minimum of maximums of each row/col
		3. keeping track of the board taking minimum and maximum number of moves to "Bingo!"
		4. score only the min/max boards by summing tiles with a move index >= number of moves required to "Bingo!"

		In total this performs ~15-20x, excluding reading in the input moves and boards.
	*/

	inputs := util.ReadLines("inputs/day4.txt")
	moves := strings.Split(inputs[0], ",")
	boards := parseBoards(inputs[1:])

	// remap moves to the turn in which that move is called
	moveToCall := map[string]int{}
	for i, m := range moves {
		moveToCall[m] = i
	}

	// remap boards to called spaced
	for _, board := range boards {
		board.tilesd = make([][]int, 5)
		for i := 0; i < 5; i++ {
			board.tilesd[i] = make([]int, 5)
			for j := 0; j < 5; j++ {
				board.tilesd[i][j] = moveToCall[board.tiles[i][j]]
			}
		}
	}

	minBoard := &board{
		moves: 100,
	}
	maxBoard := &board{
		moves: 1,
	}

	// go through each board and find the number of moves it takes to win
	for _, board := range boards {
		minMovesToBingo := 100
		for i := 0; i < 5; i++ {
			maxMovesPerRow := 0
			maxMovesPerCol := 0
			for j := 0; j < 5; j++ {
				moveByRow := board.tilesd[i][j]
				if moveByRow > maxMovesPerRow {
					maxMovesPerRow = moveByRow
				}
				moveByCol := board.tilesd[j][i]
				if moveByCol > maxMovesPerCol {
					maxMovesPerCol = moveByCol
				}
			}
			if maxMovesPerCol < minMovesToBingo {
				minMovesToBingo = maxMovesPerCol
			}
			if maxMovesPerRow < minMovesToBingo {
				minMovesToBingo = maxMovesPerRow
			}
		}
		board.moves = minMovesToBingo
		if board.moves > maxBoard.moves {
			maxBoard = board
		}
		if board.moves < minBoard.moves {
			minBoard = board
		}
	}

	score := func(b *board) {
		score := 0
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if b.tilesd[i][j] > b.moves {
					score += util.ParseInt(b.tiles[i][j])
				}
			}
		}
		finalMove := moves[b.moves]
		fmt.Printf("Bingo on move %s with a subscore of %d and total score of %d\n", finalMove, score, score*util.ParseInt(finalMove))
	}
	fmt.Println("Min board:", minBoard.moves)
	score(minBoard)
	fmt.Println("Max board:", maxBoard.moves)
	score(maxBoard)
}

func basicSolution() {
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
	n      int
	moves  int
	tiles  [][]string
	tilesd [][]int
}

func (b *board) print() {
	fmt.Println("Board", b.n)
	for _, row := range b.tiles {
		fmt.Println(row)
	}
}

func (b *board) printd() {
	fmt.Println("Board", b.n)
	for _, row := range b.tilesd {
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
		boardRawLines := inputs[i+1 : i+6]
		b := &board{
			n:     len(boards),
			tiles: make([][]string, 5)[:0],
		}
		for _, line := range boardRawLines {
			row := make([]string, 5)[:0]
			for j := 0; j < 14; j += 3 {
				row = append(row, strings.TrimSpace(line[j:j+2]))
			}
			b.tiles = append(b.tiles, row)
		}
		boards = append(boards, b)
	}
	return boards
}
