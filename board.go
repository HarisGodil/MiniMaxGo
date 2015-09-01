package main

// Board is a generic that can be used for many different games (with TicTacToe in mind)
type Board interface {
	possibleMoves() ([]Move, error)
	makeMove(a Move) (Board, error)
	whoseTurn() bool
	checkForWin() player
	printBoard() string
	isMoveValid(a Move) bool
}

// Move is used for potential ai moves as well as human moves
type Move struct {
	x    int
	y    int
	mark player
}

// -1 for player 0, 0 for empty, 1 for player 1
type player int
