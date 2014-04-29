package main

import (
	"fmt"

	"github.com/alemedeiros/nnc"
)

// printBoard prints board to stdin.
func printBoard(board [][]byte) {
	str := make([]byte, 0)
	for i, l := range board {

		// Print board lines
		if i != 0 {
			for j, _ := range l {
				if j == 0 {
					str = append(str, '-')
				} else {
					str = append(str, '+', '-')
				}
			}
			str = append(str, '\n')
		}

		// Print elements
		for j, x := range l {
			if j == 0 {
				str = append(str, x)
			} else {
				str = append(str, '|', x)
			}
		}
		str = append(str, '\n')
	}

	fmt.Printf("%s", str)
}

func main() {
	var g nnc.Game
	var n int
	var win byte

	// Get board size from input and instantiates new game.
	fmt.Print("Size of the board: ")
	fmt.Scan(&n)
	g = nnc.New(n)

	for {
		var i, j int
		var end bool
		var err error

		// Print board status
		printBoard(g.Board())

		// Outcome functions for debug purpose only
		fmt.Println("Outcome(X):", g.Outcome(nnc.Cross))
		fmt.Println("Outcome(O):", g.Outcome(nnc.Nought))

		// Get coordinates from player
		curr := g.CurrentPlayer()
		fmt.Printf("Player %c enter your coordinates: ", curr)
		fmt.Scan(&i, &j)
		fmt.Println()

		end, win, err = g.Play(i, j, curr)

		// Check for errors
		if err != nil {
			// TODO: Print error to stderr
			fmt.Printf("nnc: %s\n", err)
			fmt.Println("Please try again.")
			continue
		}

		// Verify if game ended
		if end {
			break
		}
	}

	// Print final board status.
	printBoard(g.Board())

	// Show winner
	switch win {
	case nnc.Empty:
		fmt.Println("Draw!")
	default:
		fmt.Printf("Player %c won!\n", win)
	}
}
