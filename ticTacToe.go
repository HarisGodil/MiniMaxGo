package main

import (
	"fmt"
)

// TicTacToeBoard is a implementation of Board
type TicTacToeBoard struct {
	board [][]player
}

// NewTicTacToeBoard creates an empty TicTacToeBoard
func NewTicTacToeBoard() TicTacToeBoard {
	return TicTacToeBoard{
		board: [][]player{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
	}
}

func (t TicTacToeBoard) possibleMoves() []Move {
	moves := make([]Move, 0, 9)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			move := Move{x: i, y: j, mark: ai}
			if t.isMoveValid(move) {
				moves = append(moves, move)
			}

		}
	}

	return moves
}

func (t TicTacToeBoard) makeMove(a Move) (Board, error) {
	if !t.isMoveValid(a) {
		return nil, fmt.Errorf("INVALID MOVE")
	}

	return updateBoard(t.board, a), nil
}

func updateBoard(oldBoard [][]player, a Move) TicTacToeBoard {
	updated := NewTicTacToeBoard()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == a.x && j == a.y {
				updated.board[i][j] = a.mark
			} else {
				updated.board[i][j] = oldBoard[i][j]
			}
		}
	}

	return updated
}

func (t TicTacToeBoard) isHumanTurn() bool {

	total := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			total += int(t.board[i][j])
		}
	}

	return total == 0
}

func (t TicTacToeBoard) checkForWin() player {
	total := 0

	// check for row victories
	for i := 0; i < 3; i++ {
		total = 0

		for j := 0; j < 3; j++ {
			total += int(t.board[i][j])
		}

		if total == 3 || total == -3 {
			return player(total / 3)
		}
	}

	// check for column victories
	for j := 0; j < 3; j++ {
		total = 0

		for i := 0; i < 3; i++ {
			total += int(t.board[i][j])
		}

		if total == 3 || total == -3 {
			return player(total / 3)
		}
	}

	// check for diagonal \ victories
	total = 0
	for x := 0; x < 3; x++ {
		total += int(t.board[x][x])
	}
	if total == 3 || total == -3 {
		return player(total / 3)
	}

	// check for diagonal / victories
	total = 0
	for x := 0; x < 3; x++ {
		total += int(t.board[2-x][x])
	}
	if total == 3 || total == -3 {
		return player(total / 3)
	}

	return player(0)
}

func (t TicTacToeBoard) printBoard() string {
	board := ""
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			board = fmt.Sprintf("%s %s ", board, intToChar(t.board[i][j]))
		}
		board += "\n"
	}
	return board
}

func (t TicTacToeBoard) isMoveValid(a Move) bool {

	if !isProperIndex(a.x) {
		return false
	}
	if !isProperIndex(a.y) {
		return false
	}

	if t.board[a.x][a.y] == 0 {
		return true
	}

	return false
}

func isProperIndex(val int) bool {
	return val >= 0 && val <= 2
}

func intToChar(val player) string {
	switch val {
	case 1:
		return "A"
	case -1:
		return "H"
	case 0:
		return "0"
	}

	return "."
}
