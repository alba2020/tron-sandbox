package main

import "fmt"

// ------------ start player -------------

type Player struct {
	id       int8
	x        int8
	y        int8
	world    *World
	nextTurn AI
}

func (p *Player) NextTurn() Direction {
	return p.nextTurn(p.world, p)
}

func (p *Player) NextCoords(turn Direction) (int8, int8) {
	switch turn {
	case Left:
		return p.x - 1, p.y
	case Right:
		return p.x + 1, p.y
	case Up:
		return p.x, p.y - 1
	case Down:
		return p.x, p.y + 1
	}

	panic(fmt.Sprintf("bad turn: [%s]", turn))
}

// ----------- end player ------------
