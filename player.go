package main

import "fmt"

type Player struct {
	id       int8
	x        int8
	y        int8
	world    *World
	nextTurn AI
}

func (p *Player) NextTurn() string {
	return p.nextTurn(p.world, p)
}

func (p *Player) NextCoords(turn string) (int8, int8) {
	switch turn {
	case "LEFT":
		return p.x - 1, p.y
	case "RIGHT":
		return p.x + 1, p.y
	case "UP":
		return p.x, p.y - 1
	case "DOWN":
		return p.x, p.y + 1
	}
	fmt.Println("bad turn")
	return p.x, p.y
}
