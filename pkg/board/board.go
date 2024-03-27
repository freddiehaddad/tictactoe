// Package board represents the Tic Tac Toe game board.
// A 3x3 grid that players alternate drawing marks on
// until the game ends.
package board

import (
	"fmt"

	"github.com/freddiehaddad/tictactoe/pkg/mark"
)

const (
	COLUMNS = 3
	ROWS    = 3
	CELLS   = ROWS * COLUMNS

	TOP_LEFT     = '┏'
	TOP_RIGHT    = '┓'
	BOTTOM_LEFT  = '┗'
	BOTTOM_RIGHT = '┛'

	TOP_SEPARATOR    = '┯'
	BOTTOM_SEPARATOR = '┷'
	LEFT_SEPARATOR   = '┠'
	RIGHT_SEPARATOR  = '┨'

	TOP_BORDER    = '━'
	BOTTOM_BORDER = '━'
	LEFT_BORDER   = '┃'
	RIGHT_BORDER  = '┃'

	INSIDE_CROSS      = '┼'
	INSIDE_HORIZONTAL = '─'
	INSIDE_VERTICAL   = '│'
)

type Board struct {
	Board []*mark.Mark
}

var CellMap map[string]int

func init() {
	CellMap = map[string]int{
		"A0": 0,
		"A1": 1,
		"A2": 2,
		"B0": 3,
		"B1": 4,
		"B2": 5,
		"C0": 6,
		"C1": 7,
		"C2": 8,
	}
}

func New() *Board {
	board := &Board{}

	for cell := 0; cell < CELLS; cell++ {
		board.Board = append(board.Board, mark.New())
	}

	return board
}

func (b *Board) Reset() {
	for cell := 0; cell < CELLS; cell++ {
		b.Board[cell].Reset()
	}
}

func (b *Board) SetMark(index, state int) bool {
	if b.Board[index].State() != mark.OPEN {
		return false
	}

	switch state {
	case mark.X:
		b.Board[index].X()
		return true
	case mark.O:
		b.Board[index].O()
		return true
	}
	return false
}

func (b *Board) Draw() {
	fmt.Printf("     %d   %d   %d\n", 0, 1, 2)
	fmt.Printf("   %c", TOP_LEFT)
	fmt.Printf("%c%c%c", TOP_BORDER, TOP_BORDER, TOP_BORDER)
	fmt.Printf("%c", TOP_SEPARATOR)
	fmt.Printf("%c%c%c", TOP_BORDER, TOP_BORDER, TOP_BORDER)
	fmt.Printf("%c", TOP_SEPARATOR)
	fmt.Printf("%c%c%c", TOP_BORDER, TOP_BORDER, TOP_BORDER)
	fmt.Printf("%c", TOP_RIGHT)
	fmt.Printf("%c", '\n')

	fmt.Printf(" %c ", 'A')
	fmt.Printf("%c", LEFT_BORDER)
	fmt.Printf(" %c ", b.Board[0].Draw())
	fmt.Printf("%c", INSIDE_VERTICAL)
	fmt.Printf(" %c ", b.Board[1].Draw())
	fmt.Printf("%c", INSIDE_VERTICAL)
	fmt.Printf(" %c ", b.Board[2].Draw())
	fmt.Printf("%c", RIGHT_BORDER)
	fmt.Printf("%c", '\n')

	fmt.Printf("   %c", LEFT_SEPARATOR)
	fmt.Printf("%c%c%c", INSIDE_HORIZONTAL, INSIDE_HORIZONTAL, INSIDE_HORIZONTAL)
	fmt.Printf("%c", INSIDE_CROSS)
	fmt.Printf("%c%c%c", INSIDE_HORIZONTAL, INSIDE_HORIZONTAL, INSIDE_HORIZONTAL)
	fmt.Printf("%c", INSIDE_CROSS)
	fmt.Printf("%c%c%c", INSIDE_HORIZONTAL, INSIDE_HORIZONTAL, INSIDE_HORIZONTAL)
	fmt.Printf("%c", RIGHT_SEPARATOR)
	fmt.Printf("%c", '\n')

	fmt.Printf(" %c ", 'B')
	fmt.Printf("%c", LEFT_BORDER)
	fmt.Printf(" %c ", b.Board[3].Draw())
	fmt.Printf("%c", INSIDE_VERTICAL)
	fmt.Printf(" %c ", b.Board[4].Draw())
	fmt.Printf("%c", INSIDE_VERTICAL)
	fmt.Printf(" %c ", b.Board[5].Draw())
	fmt.Printf("%c", RIGHT_BORDER)
	fmt.Printf("%c", '\n')

	fmt.Printf("   %c", LEFT_SEPARATOR)
	fmt.Printf("%c%c%c", INSIDE_HORIZONTAL, INSIDE_HORIZONTAL, INSIDE_HORIZONTAL)
	fmt.Printf("%c", INSIDE_CROSS)
	fmt.Printf("%c%c%c", INSIDE_HORIZONTAL, INSIDE_HORIZONTAL, INSIDE_HORIZONTAL)
	fmt.Printf("%c", INSIDE_CROSS)
	fmt.Printf("%c%c%c", INSIDE_HORIZONTAL, INSIDE_HORIZONTAL, INSIDE_HORIZONTAL)
	fmt.Printf("%c", RIGHT_SEPARATOR)
	fmt.Printf("%c", '\n')

	fmt.Printf(" %c ", 'C')
	fmt.Printf("%c", LEFT_BORDER)
	fmt.Printf(" %c ", b.Board[6].Draw())
	fmt.Printf("%c", INSIDE_VERTICAL)
	fmt.Printf(" %c ", b.Board[7].Draw())
	fmt.Printf("%c", INSIDE_VERTICAL)
	fmt.Printf(" %c ", b.Board[8].Draw())
	fmt.Printf("%c", RIGHT_BORDER)
	fmt.Printf("%c", '\n')

	fmt.Printf("   %c", BOTTOM_LEFT)
	fmt.Printf("%c%c%c", BOTTOM_BORDER, BOTTOM_BORDER, BOTTOM_BORDER)
	fmt.Printf("%c", BOTTOM_SEPARATOR)
	fmt.Printf("%c%c%c", BOTTOM_BORDER, BOTTOM_BORDER, BOTTOM_BORDER)
	fmt.Printf("%c", BOTTOM_SEPARATOR)
	fmt.Printf("%c%c%c", BOTTOM_BORDER, BOTTOM_BORDER, BOTTOM_BORDER)
	fmt.Printf("%c", BOTTOM_RIGHT)
	fmt.Printf("%c", '\n')
}
