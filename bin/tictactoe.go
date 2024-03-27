package main

import (
	"fmt"

	"github.com/freddiehaddad/tictactoe/pkg/game"
)

func main() {
	fmt.Println("Welcome to Tic Tac Toe!")
	game := game.New()
	game.Start()
	fmt.Println("\nExiting!")
}
