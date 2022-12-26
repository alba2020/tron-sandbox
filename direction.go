package main

// --------- start direction -------

type Direction int8

const (
	Left Direction = iota
	Up
	Right
	Down
)

func (d Direction) String() string {
	switch d {
	case Left:
		return "LEFT"
	case Up:
		return "UP"
	case Right:
		return "RIGHT"
	case Down:
		return "DOWN"
	}
	return "unknown"
}

var DIRECTIONS = []Direction{
	Left, Up, Right, Down,
}

// ------ end direction ------------
