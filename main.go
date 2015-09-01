package main

import (
	"fmt"
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
	fmt.Print("TicTacToe")
	isBoard(NewTicTacToeBoard())
}

func isBoard(b Board) {
	fmt.Printf(" IS A BOARD\n")
	fmt.Printf("%s", b.printBoard())
}
