package main

import (
	"flag"
	"fmt"

	"github.com/alemedeiros/nnc"
)

// Flag variables
var human bool
var serial bool

// Set command line flags
func initFlags() {
	flag.BoolVar(&human, "human", false, "Human player for Cross")
	flag.BoolVar(&serial, "serial", false, "Serial AI functions are used instead")
}

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

	initFlags()
	flag.Parse()

	// Get board size from input and instantiates new game.
	fmt.Print("Size of the board: ")
	fmt.Scan(&n)
	g = nnc.New(n)

	for {
		var end bool
		var err error
		var curr byte

		// Print board status
		printBoard(g.Board())

		curr = g.CurrentPlayer()
		fmt.Println()

		if curr == nnc.Cross {
			if human {
				var i, j int

				// Get Human player move
				fmt.Printf("Player %c enter your coordinates: ", curr)
				fmt.Scanln(&i, &j)
				fmt.Println()

				end, win, err = g.Play(i, j, curr)
			} else {
				// AI Player
				if serial {
					end, win, err = g.PlayAISerial(curr)
				} else {
					end, win, err = g.PlayAI(curr)
				}
			}
		} else {
			// Nought is always AI Player
			if serial {
				end, win, err = g.PlayAISerial(curr)
			} else {
				end, win, err = g.PlayAI(curr)
			}
		}

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
