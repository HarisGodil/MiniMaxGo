package main

import (
	"fmt"
	"math/rand"
)

type heuristicAction func(Board) int

// BasicTicTacToeHeuristic is a heuristicAction that just relies on win conditions
func BasicTicTacToeHeuristic(board Board) int {
	return int(board.checkForWin())
}

// DecideMove will be given a heuristicAction, that searches till runTime is up
func DecideMove(action heuristicAction, b Board, runTime int) (Move, error) {

	if b.isHumanTurn() {
		return Move{}, fmt.Errorf("It is not the AI's turn")
	}

	//var bestMove Move

	moves := b.possibleMoves()
	index := rand.Int31n(int32(len(moves)))

	return moves[index], nil

	//return bestMove, nil
}
