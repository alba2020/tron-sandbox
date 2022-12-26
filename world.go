package main

import (
	"fmt"
	"math/rand"
)

// --------- start world ---------

const EMPTY int8 = -1
const NUM_PLAYERS = 4

type World struct {
	Grid
	players [NUM_PLAYERS]*Player
	alive   int8
	current int8
}

func NewWorld(w, h int8) World {
	cells := make([]int8, int(w)*int(h))
	for i := range cells {
		cells[i] = EMPTY
	}

	return World{
		Grid:    Grid{w, h, cells},
		players: [NUM_PLAYERS]*Player{},
		alive:   0,
	}
}

func ParseWorld(w int8, h int8, txt string) *World {
	return &World{
		Grid: ParseGrid(w, h, txt),
	}
}

func (w *World) Copy() *World {
	cp := World{
		Grid:    w.Grid.Copy(),
		players: [NUM_PLAYERS]*Player{},
		alive:   w.alive,
	}

	for i := 0; i < NUM_PLAYERS; i++ {
		if w.players[i] == nil {
			cp.players[i] = nil
		} else {
			cp.players[i] = &Player{
				id:       w.players[i].id,
				x:        w.players[i].x,
				y:        w.players[i].y,
				world:    &cp,
				nextTurn: w.players[i].nextTurn,
			}
		}
	}

	return &cp
}

func (w *World) AddPlayer(id, x, y int8, f AI) *Player {
	p := Player{
		id:       id,
		x:        x,
		y:        y,
		world:    w,
		nextTurn: f,
	}
	w.players[id] = &p
	w.alive++
	w.Set(p.x, p.y, p.id)
	return &p
}

func (w *World) NextPlayer(prev int8) int8 {
	for i := prev + 1; i <= prev+NUM_PLAYERS; i++ {
		id := i % NUM_PLAYERS
		if w.players[id] != nil {
			return id
		}
	}
	return -1
}

func (w *World) AddPlayerAtRandom(id int8, f AI) {
	rx := int8(rand.Intn(int(w.width)))
	ry := int8(rand.Intn(int(w.height)))
	w.AddPlayer(id, rx, ry, f)
}

func (w *World) ValidTurns(p *Player) []Direction {
	// valid := make([]string, 0, 4)
	valid := []Direction{}

	for _, turn := range DIRECTIONS {
		if w.TurnValid(p.NextCoords(turn)) {
			valid = append(valid, turn)
		}
	}

	return valid
}

func (w *World) ApplyTurn(pid int8, turn Direction) {
	player := w.players[pid]
	player.x, player.y = player.NextCoords(turn)

	if player.x < -1 || player.y < -1 {
		panic(fmt.Sprintf("WHAT THE FUCK!!! %v", player))
	}

	if !w.TurnValid(player.x, player.y) {

		// fmt.Printf("player %d turn %s invalid\n", player.id, turn)
		// fmt.Println("player dump", player)
		// fmt.Println(w)

		w.players[player.id] = nil
		w.alive--
		w.ClearPath(player.id)
	} else {
		w.Set(player.x, player.y, player.id)
	}
}

func (w *World) Active() bool {
	return w.alive > 1
}

func (w *World) SimTurn() {
	player := w.players[w.current]
	turn := player.NextTurn()

	// if log {
	// 	fmt.Printf("[%d] [%s]\n", player.id, turn)
	// }

	w.ApplyTurn(w.current, turn) // 2
	w.current = w.NextPlayer(w.current)
}

func (w *World) Simulate() int8 {
	for w.Active() {
		w.SimTurn()
	}
	winner := w.GetWinner()
	return winner.id
}

func (w *World) GetWinner() *Player {
	if w.alive != 1 {
		panic("should be 1 winner")
	}
	for i := range w.players {
		if w.players[i] != nil {
			return w.players[i]
		}
	}
	return nil
}

// ---------- end world -------------
