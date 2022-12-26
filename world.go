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
	players [NUM_PLAYERS]Player
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
		players: [NUM_PLAYERS]Player{},
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
		players: [NUM_PLAYERS]Player{},
		alive:   w.alive,
	}

	for i := 0; i < NUM_PLAYERS; i++ {
		if w.players[i].alive {
			cp.players[i] = w.players[i]
			cp.players[i].world = &cp
		}
	}

	return &cp
}

func (w *World) AddPlayer(id, x, y int8, f AI) *Player {
	w.players[id] = Player{
		id:       id,
		x:        x,
		y:        y,
		world:    w,
		nextTurn: f,
		alive:    true,
	}
	w.alive++
	w.Set(x, y, id)
	return &w.players[id]
}

func (w *World) NextPlayer(prev int8) int8 {
	for i := prev + 1; i <= prev+NUM_PLAYERS; i++ {
		id := i % NUM_PLAYERS
		if w.players[id].alive {
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

func (w *World) ValidTurns(pid int8) DirectionMask {
	valid := None

	p := &w.players[pid]

	for _, direction := range DIRECTIONS {
		if w.TurnValid(p.NextCoords(direction)) {
			valid |= direction // set bit
		}
	}

	return valid
}

func (w *World) ApplyTurn(pid int8, turn DirectionMask) {
	player := &w.players[pid]
	player.x, player.y = player.NextCoords(turn)

	if player.x < -1 || player.y < -1 {
		panic(fmt.Sprintf("WHAT THE FUCK!!! %v", player))
	}

	if !w.TurnValid(player.x, player.y) {

		// fmt.Printf("player %d turn %s invalid\n", player.id, turn)
		// fmt.Println("player dump", player)
		// fmt.Println(w)

		w.players[player.id].alive = false
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
	player := &w.players[w.current]
	turn := player.NextTurn()

	// fmt.Printf("%v [%s]\n", player, turn)

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
		if w.players[i].alive {
			return &w.players[i]
		}
	}
	return nil
}

// ---------- end world -------------
