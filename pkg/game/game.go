package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/freddiehaddad/tictactoe/pkg/board"
	"github.com/freddiehaddad/tictactoe/pkg/mark"
)

type Game struct {
	Board *board.Board
	state int
}

func New() *Game {
	return &Game{
		Board: board.New(),
	}
}

func (g *Game) Turn() byte {
	switch g.state {
	case mark.X:
		return byte(mark.X)
	case mark.O:
		return byte(mark.O)
	default:
		fmt.Fprintf(os.Stderr, "Turn error: Invalid state %d", g.state)
		return 0
	}
}

func (g *Game) ToggleTurn() {
	switch g.state {
	case mark.X:
		g.state = mark.O
	case mark.O:
		g.state = mark.X
	default:
		fmt.Fprintf(os.Stderr, "ToggleTurn error: Invalid state %d", g.state)
	}
}

func (g *Game) CheckOver() bool {
	for cell := 0; cell < board.CELLS; cell++ {
		if g.Board.Board[cell].State() == mark.OPEN {
			return false
		}
	}
	return true
}

func (g *Game) CheckWinner() bool {
	// Top Horizontal
	if g.Board.Board[0].State() != mark.OPEN {
		if g.Board.Board[0].State() == g.Board.Board[1].State() && g.Board.Board[0].State() == g.Board.Board[2].State() {
			g.Board.Board[3].Reset()
			g.Board.Board[4].Reset()
			g.Board.Board[5].Reset()
			g.Board.Board[6].Reset()
			g.Board.Board[7].Reset()
			g.Board.Board[8].Reset()
			return true
		}
	}

	// Middle Horizontal
	if g.Board.Board[3].State() != mark.OPEN {
		if g.Board.Board[3].State() == g.Board.Board[4].State() && g.Board.Board[3].State() == g.Board.Board[5].State() {
			g.Board.Board[0].Reset()
			g.Board.Board[1].Reset()
			g.Board.Board[2].Reset()
			g.Board.Board[6].Reset()
			g.Board.Board[7].Reset()
			g.Board.Board[8].Reset()
			return true
		}
	}

	// Bottom Horizontal
	if g.Board.Board[6].State() != mark.OPEN {
		if g.Board.Board[6].State() == g.Board.Board[7].State() && g.Board.Board[6].State() == g.Board.Board[8].State() {
			g.Board.Board[0].Reset()
			g.Board.Board[1].Reset()
			g.Board.Board[2].Reset()
			g.Board.Board[3].Reset()
			g.Board.Board[4].Reset()
			g.Board.Board[5].Reset()
			return true
		}
	}

	// Left Vertical
	if g.Board.Board[0].State() != mark.OPEN {
		if g.Board.Board[0].State() == g.Board.Board[3].State() && g.Board.Board[0].State() == g.Board.Board[6].State() {
			g.Board.Board[1].Reset()
			g.Board.Board[2].Reset()
			g.Board.Board[4].Reset()
			g.Board.Board[5].Reset()
			g.Board.Board[7].Reset()
			g.Board.Board[8].Reset()
			return true
		}
	}

	// Middle Vertical
	if g.Board.Board[1].State() != mark.OPEN {
		if g.Board.Board[1].State() == g.Board.Board[4].State() && g.Board.Board[1].State() == g.Board.Board[7].State() {
			g.Board.Board[0].Reset()
			g.Board.Board[2].Reset()
			g.Board.Board[3].Reset()
			g.Board.Board[5].Reset()
			g.Board.Board[6].Reset()
			g.Board.Board[8].Reset()
			return true
		}
	}

	// Right Vertical
	if g.Board.Board[2].State() != mark.OPEN {
		if g.Board.Board[2].State() == g.Board.Board[5].State() && g.Board.Board[2].State() == g.Board.Board[8].State() {
			g.Board.Board[0].Reset()
			g.Board.Board[1].Reset()
			g.Board.Board[3].Reset()
			g.Board.Board[4].Reset()
			g.Board.Board[6].Reset()
			g.Board.Board[7].Reset()
			return true
		}
	}

	// Top left to bottom right diagnonal
	if g.Board.Board[0].State() != mark.OPEN {
		if g.Board.Board[0].State() == g.Board.Board[4].State() && g.Board.Board[0].State() == g.Board.Board[8].State() {
			g.Board.Board[1].Reset()
			g.Board.Board[2].Reset()
			g.Board.Board[3].Reset()
			g.Board.Board[5].Reset()
			g.Board.Board[6].Reset()
			g.Board.Board[7].Reset()
			return true
		}
	}

	// Bottom left to top right diagnonal
	if g.Board.Board[6].State() != mark.OPEN {
		if g.Board.Board[6].State() == g.Board.Board[4].State() && g.Board.Board[6].State() == g.Board.Board[2].State() {
			g.Board.Board[0].Reset()
			g.Board.Board[1].Reset()
			g.Board.Board[3].Reset()
			g.Board.Board[5].Reset()
			g.Board.Board[7].Reset()
			g.Board.Board[8].Reset()
			return true
		}
	}

	return false
}

func (g *Game) Start() {
	scanner := bufio.NewScanner(os.Stdin)

	g.Board.Reset()
	g.state = mark.X
	for {
		g.Board.Draw()
		fmt.Printf("%c: Type a row and column for your mark (e.g. B2): ", g.Turn())
		if !scanner.Scan() {
			return
		}

		input := strings.ToUpper(scanner.Text())
		switch input {
		case "Q":
			return
		default:
			if index, ok := board.CellMap[input]; !ok {
				fmt.Printf("Invalid input %q\n", input)
				continue
			} else if !g.Board.SetMark(index, int(g.Turn())) {
				fmt.Printf("%q has already been marked with %q\n", input, g.Board.Board[index].Draw())
				continue
			}

		}

		if g.CheckWinner() {
			g.Board.Draw()
			fmt.Printf("%c wins!\n", g.Turn())
			g.Board.Reset()
			continue
		}

		if g.CheckOver() {
			fmt.Printf("Tie game!\n")
			g.Board.Reset()
			continue
		}

		g.ToggleTurn()
	}
}
