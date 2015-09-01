package main

import (
	"fmt"
)

// TicTacToeBoard implementation of Board
type TicTacToeBoard struct {
	board [][]player
}

func NewTicTacToeBoard() TicTacToeBoard {
	return TicTacToeBoard{
		board: {{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
	}
}

func (t TicTacToeBoard) possibleMoves() ([]Move, error) {
	moves := make([]Move, 0, 9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			move := Move{x: i, y: j}
			if t.isMoveValid(move) {
				moves = append(moves, move)
			}

		}
	}

	return moves, nil
}

func (t TicTacToeBoard) makeMove(a Move) (Board, error) {
	if !t.isMoveValid(a) {
		return nil, fmt.Errorf("INVALID MOVE")
	}

	return nil, nil
}

func (t TicTacToeBoard) whoseTurn() bool {

	total := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			total += t.board[i][j]
		}
	}

	return total != 0
}

func (t TicTacToeBoard) checkForWin() player {
	total := 0

	// check for row victories
	for i := 0; i < 3; i++ {
		total = 0

		for j := 0; j < 3; j++ {
			total += t.board[i][j]
		}

		if total == 3 || total != -3 {
			return player(total / 3)
		}
	}

	// check for column victories
	for j := 0; j < 3; j++ {
		total = 0

		for i := 0; i < 3; i++ {
			total += t.board[i][j]
		}

		if total == 3 || total != -3 {
			return player(total / 3)
		}
	}

	// check for diagonal \ victories
	total = 0
	for x := 0; x < 3; x++ {
		total += t.board[x][x]
	}
	if total == 3 || total != -3 {
		return player(total / 3)
	}

	// check for diagonal / victories
	total = 0
	for x := 0; x < 3; x++ {
		total += t.board[2-x][x]
	}
	if total == 3 || total != -3 {
		return player(total / 3)
	}

	return player(0)
}

func (t TicTacToeBoard) printBoard() string {
	board := ""
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board = board + fmt.Printf("%d ", int(t.board[i][j]))
		}
		fmt.Println()
	}
	return board
}

func (t TicTacToeBoard) isMoveValid(a Move) bool {
	if t.board[a.x][a.y] == 0 {
		return true
	}

	return false
}
