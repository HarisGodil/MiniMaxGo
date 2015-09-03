package main

import (
	"fmt"
	"log"
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

	move, _, done := BestMove(action, b, 4)
	done = done
	return move, nil

	/*moves := b.possibleMoves()
	index := rand.Int31n(int32(len(moves)))
	fmt.Printf("%d\n", index)

	return moves[index], nil*/
}

// BestMove takes in a heuristicAction (specific to the board implementations), and board and depth to go further
// Returns the best move to make, score that it would result in, whether it needs to search deeper
func BestMove(action heuristicAction, b Board, depth int) (Move, int, bool) {

	if win := b.checkForWin(); win == 0 {

		// somehow is the best move......
		bestMove := Move{-3, -3, ai}
		bestScore = 0
		var done bool

		if depth == 1 {

			moves := b.possibleMoves()
			for _, move := range moves {

				tmpBoard, err := b.makeMove(move)
				if err != nil {
					log.Fatal(err)
				}

				multiplier := human
				if tmpBoard.isHumanTurn() { // meaning that the ai just made a move
					multiplier = ai
				}

				score := action(tmpBoard)

				if score*multiplier > bestScore*multiplier {
					bestScore = score
					bestMove = move
					done = tmpBoard.checkForWin() == 0
				}

			}

		} else {
			for _, move := range b.possibleMoves() {

				tmpBoard, err := b.makeMove(move)
				if err != nil {
					log.Fatal(err)
				}

				multiplier := human
				if tmpBoard.isHumanTurn() { // meaning that the ai just made a move
					multiplier = ai
				}

				// move is ignored because it isnt the most immediate move
				_, score, isDone := BestMove(action, tmpBoard, depth-1)

				if score*multiplier > bestScore*multiplier {
					bestScore = score
					bestMove = move
					done = isDone
				}

			}
		}
		return bestMove, bestScore, done

	}
	log.Fatalf("%d %d\n%s", depth, int(b.checkForWin()), b.printBoard())
	// this method shouldn't be called on a board that is over
	// so it returns an invalid move (recursive calls only need the score)
	return Move{-2, -2, ai}, int(b.checkForWin()), true
}
