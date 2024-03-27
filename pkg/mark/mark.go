// Package mark represents the X and O marks used in Tic Tac Toe
package mark

const (
	OPEN = 0
	X        = int('X')
	O        = int('O')
)

type Mark struct {
	state int
}

func New() *Mark {
	return &Mark{state: OPEN}
}

func (p *Mark) Reset() {
	p.state = OPEN
}

func (p *Mark) State() int {
	return p.state
}

func (p *Mark) X() {
	p.state = X
}

func (p *Mark) O() {
	p.state = O
}

func (p *Mark) Draw() byte {
	switch p.state {
	case X:
		return 'X'
	case O:
		return 'O'
	default:
		return ' '
	}
}
