package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var runTime = cli.IntFlag{
	Name:  "runTime",
	Value: 1,
	Usage: "number of seconds that the AI will have to look ahead",
}

func main() {
	app := cli.NewApp()
	app.Name = "minimax"
	app.Usage = "to play against a computer using a minimax algorithm"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{ticTacToeCmd}
	app.Run(os.Args)
}

var ticTacToeCmd = cli.Command{
	Name:    "tictactoe",
	Usage:   "play tictactoe against an AI",
	Aliases: []string{"m"},
	Flags:   []cli.Flag{runTime},
	Action:  runTicTacToe,
}

func runTicTacToe(c *cli.Context) {
	b := NewTicTacToeBoard()

	for {

		fmt.Printf("%s", b.printBoard())

		if win := b.checkForWin(); win != 0 {
			if win == human {
				fmt.Print("Player Wins\n")
				return
			}
			fmt.Print("AI Wins\n")
			return
		}

		if b.isHumanTurn() {

			fmt.Print("Player Turn\n")
			move := humanMove(b)

			tempBoard, err := b.makeMove(move)
			if err != nil {
				fmt.Errorf("Error in making turn: |%v|", err)
			}

			b = tempBoard.(TicTacToeBoard)

		} else {

			fmt.Print("AI Turn\nCalculating move...\n")

			move, err := DecideMove(BasicTicTacToeHeuristic, b, c.Int("runTime"))
			if err != nil {
				fmt.Errorf("Error in deciding move: |%v|", err)
			}

			tempBoard, err := b.makeMove(move)
			if err != nil {
				fmt.Errorf("Error in making turn: |%v|", err)
				log.Fatalf("WTFBOOM %d %d", int(move.x), int(move.y))
			}

			fmt.Printf("err: %v", err)

			b = tempBoard.(TicTacToeBoard)
		}
	}
}

func humanMove(b Board) Move {
	var x, y int
	fmt.Scanf("%d %d", &x, &y)

	move := Move{x - 1, y - 1, human}
	if !b.isMoveValid(move) {
		fmt.Print("That move was invalid\n")
		return humanMove(b)
	}

	return move
}
