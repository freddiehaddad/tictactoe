// Package board represents the Tic Tac Toe game board.
// A 3x3 grid that players alternate drawing marks on
// until the game ends.
package board

import (
	"fmt"
	"strings"

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
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("     %d   %d   %d\n", 0, 1, 2))
	sb.WriteString(fmt.Sprintf("   %c", TOP_LEFT))
	sb.WriteString(fmt.Sprintf("%c%c%c", TOP_BORDER, TOP_BORDER, TOP_BORDER))
	sb.WriteString(fmt.Sprintf("%c", TOP_SEPARATOR))
	sb.WriteString(fmt.Sprintf("%c%c%c", TOP_BORDER, TOP_BORDER, TOP_BORDER))
	sb.WriteString(fmt.Sprintf("%c", TOP_SEPARATOR))
	sb.WriteString(fmt.Sprintf("%c%c%c", TOP_BORDER, TOP_BORDER, TOP_BORDER))
	sb.WriteString(fmt.Sprintf("%c", TOP_RIGHT))
	sb.WriteString(fmt.Sprintf("%c", '\n'))

	sb.WriteString(fmt.Sprintf(" %c ", 'A'))
	sb.WriteString(fmt.Sprintf("%c", LEFT_BORDER))
	sb.WriteString(fmt.Sprintf(" %c ", b.Board[0].Draw()))
	sb.WriteString(fmt.Sprintf("%c", INSIDE_VERTICAL))
	sb.WriteString(fmt.Sprintf(" %c ", b.Board[1].Draw()))
	sb.WriteString(fmt.Sprintf("%c", INSIDE_VERTICAL))
	sb.WriteString(fmt.Sprintf(" %c ", b.Board[2].Draw()))
	sb.WriteString(fmt.Sprintf("%c", RIGHT_BORDER))
	sb.WriteString(fmt.Sprintf("%c", '\n'))

	sb.WriteString(fmt.Sprintf("   %c", LEFT_SEPARATOR))
	sb.WriteString(fmt.Sprintf("%c%c%c", INSIDE_HORIZONTAL, INSIDE_HORIZONTAL, INSIDE_HORIZONTAL))
	sb.WriteString(fmt.Sprintf("%c", INSIDE_CROSS))
	sb.WriteString(fmt.Sprintf("%c%c%c", INSIDE_HORIZONTAL, INSIDE_HORIZONTAL, INSIDE_HORIZONTAL))
	sb.WriteString(fmt.Sprintf("%c", INSIDE_CROSS))
	sb.WriteString(fmt.Sprintf("%c%c%c", INSIDE_HORIZONTAL, INSIDE_HORIZONTAL, INSIDE_HORIZONTAL))
	sb.WriteString(fmt.Sprintf("%c", RIGHT_SEPARATOR))
	sb.WriteString(fmt.Sprintf("%c", '\n'))

	sb.WriteString(fmt.Sprintf(" %c ", 'B'))
	sb.WriteString(fmt.Sprintf("%c", LEFT_BORDER))
	sb.WriteString(fmt.Sprintf(" %c ", b.Board[3].Draw()))
	sb.WriteString(fmt.Sprintf("%c", INSIDE_VERTICAL))
	sb.WriteString(fmt.Sprintf(" %c ", b.Board[4].Draw()))
	sb.WriteString(fmt.Sprintf("%c", INSIDE_VERTICAL))
	sb.WriteString(fmt.Sprintf(" %c ", b.Board[5].Draw()))
	sb.WriteString(fmt.Sprintf("%c", RIGHT_BORDER))
	sb.WriteString(fmt.Sprintf("%c", '\n'))

	sb.WriteString(fmt.Sprintf("   %c", LEFT_SEPARATOR))
	sb.WriteString(fmt.Sprintf("%c%c%c", INSIDE_HORIZONTAL, INSIDE_HORIZONTAL, INSIDE_HORIZONTAL))
	sb.WriteString(fmt.Sprintf("%c", INSIDE_CROSS))
	sb.WriteString(fmt.Sprintf("%c%c%c", INSIDE_HORIZONTAL, INSIDE_HORIZONTAL, INSIDE_HORIZONTAL))
	sb.WriteString(fmt.Sprintf("%c", INSIDE_CROSS))
	sb.WriteString(fmt.Sprintf("%c%c%c", INSIDE_HORIZONTAL, INSIDE_HORIZONTAL, INSIDE_HORIZONTAL))
	sb.WriteString(fmt.Sprintf("%c", RIGHT_SEPARATOR))
	sb.WriteString(fmt.Sprintf("%c", '\n'))

	sb.WriteString(fmt.Sprintf(" %c ", 'C'))
	sb.WriteString(fmt.Sprintf("%c", LEFT_BORDER))
	sb.WriteString(fmt.Sprintf(" %c ", b.Board[6].Draw()))
	sb.WriteString(fmt.Sprintf("%c", INSIDE_VERTICAL))
	sb.WriteString(fmt.Sprintf(" %c ", b.Board[7].Draw()))
	sb.WriteString(fmt.Sprintf("%c", INSIDE_VERTICAL))
	sb.WriteString(fmt.Sprintf(" %c ", b.Board[8].Draw()))
	sb.WriteString(fmt.Sprintf("%c", RIGHT_BORDER))
	sb.WriteString(fmt.Sprintf("%c", '\n'))

	sb.WriteString(fmt.Sprintf("   %c", BOTTOM_LEFT))
	sb.WriteString(fmt.Sprintf("%c%c%c", BOTTOM_BORDER, BOTTOM_BORDER, BOTTOM_BORDER))
	sb.WriteString(fmt.Sprintf("%c", BOTTOM_SEPARATOR))
	sb.WriteString(fmt.Sprintf("%c%c%c", BOTTOM_BORDER, BOTTOM_BORDER, BOTTOM_BORDER))
	sb.WriteString(fmt.Sprintf("%c", BOTTOM_SEPARATOR))
	sb.WriteString(fmt.Sprintf("%c%c%c", BOTTOM_BORDER, BOTTOM_BORDER, BOTTOM_BORDER))
	sb.WriteString(fmt.Sprintf("%c", BOTTOM_RIGHT))
	sb.WriteString(fmt.Sprintf("%c", '\n'))

	fmt.Print(sb.String())
}
