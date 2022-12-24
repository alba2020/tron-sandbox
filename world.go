package main

import (
	"fmt"
	"math/rand"
)

const EMPTY int8 = -1
const NUM_PLAYERS = 4

type World struct {
	width   int8
	height  int8
	cells   []int8
	players [NUM_PLAYERS]*Player
	alive   int8
	current int8
}

func (w *World) At(x int8, y int8) int8 {
	return w.cells[w.Offset(x, y)]
}

func (w *World) Set(x, y int8, val int8) {
	w.cells[w.Offset(x, y)] = val
}

func (w *World) Offset(x, y int8) int {
	return int(y)*int(w.width) + int(x)
}

func (w *World) Print() {
	var i, j int8
	for j = 0; j < w.height; j++ {
		for i = 0; i < w.width; i++ {
			if w.At(i, j) == EMPTY {
				fmt.Print(".")
			} else {
				fmt.Print(w.At(i, j))
			}
		}
		fmt.Println()
	}
}

func NewWorld(w, h int8) World {
	cells := make([]int8, int(w)*int(h))
	for i := range cells {
		cells[i] = EMPTY
	}

	return World{
		width:   w,
		height:  h,
		cells:   cells,
		players: [NUM_PLAYERS]*Player{},
		alive:   0,
	}
}

func ParseWorld(w, h int8, txt string) *World {
	cells := []int8{}

	for _, r := range txt {
		if r == '.' {
			cells = append(cells, -1)
		} else if r >= '0' && r <= '9' {
			cells = append(cells, int8(r-'0'))
		} else {
			continue
		}
	}

	if len(cells) != int(w)*int(h) {
		panic("could not parse world")
	}

	return &World{
		width:  w,
		height: h,
		cells:  cells,
	}
}

func (w *World) Copy() *World {
	cp := World{
		width:   w.width,
		height:  w.height,
		cells:   make([]int8, len(w.cells)),
		players: w.players,
		alive:   w.alive,
	}
	copy(cp.cells, w.cells)

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
	for i := prev + 1; i < prev+NUM_PLAYERS; i++ {
		id := i % NUM_PLAYERS
		if w.players[id] != nil {
			return id
		}
	}
	panic("next player not found")
	return -1
}

func (w *World) AddPlayerAtRandom(id int8, f AI) {
	rx := int8(rand.Intn(int(w.width)))
	ry := int8(rand.Intn(int(w.height)))
	w.AddPlayer(id, rx, ry, f)
}

func (w *World) GetPlayer(playerId int8) *Player {
	return w.players[playerId]
}

func (w *World) TurnValid(x, y int8) bool {
	if x < 0 || x >= w.width || y < 0 || y >= w.height {
		return false
	}

	if w.At(x, y) != EMPTY {
		return false
	}

	return true
}

func (w *World) ValidTurns(p *Player) []string {
	valid := []string{}

	for _, turn := range DIRECTIONS {
		if w.TurnValid(p.NextCoords(turn)) {
			valid = append(valid, turn)
		}
	}

	return valid
}

func (w *World) ClearPath(playerId int8) {
	for i := range w.cells {
		if w.cells[i] == playerId {
			w.cells[i] = EMPTY
		}
	}
}

func (w *World) ApplyTurn(player *Player, turn string) {
	player.x, player.y = player.NextCoords(turn)

	if !w.TurnValid(player.x, player.y) {
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

func (w *World) SimulateTurn() {
	player := w.players[w.current]
	turn := player.NextTurn()
	// fmt.Printf(" %d %s", w.players[i].id, turn)
	w.ApplyTurn(player, turn)

	if w.Active() {
		w.current = w.NextPlayer(w.current)
	}
}

func (w *World) Simulate(playerId int8) int8 {
	w.current = playerId
	for w.Active() {
		w.SimulateTurn()
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

func (w *World) Area(x, y, val int8) int {
	visited := make(map[int]struct{})

	var area func(w *World, x, y, val int8) int

	area = func(w *World, x, y, val int8) int {
		if _, found := visited[w.Offset(x, y)]; found {
			return 0
		}
		visited[w.Offset(x, y)] = struct{}{}

		if w.At(x, y) != val {
			return 0
		} else {
			sum := 0
			if y > 0 {
				sum += area(w, x, y-1, val) // up
			}
			if y < w.height-1 {
				sum += area(w, x, y+1, val) // down
			}
			if x > 0 {
				sum += area(w, x-1, y, val) // left
			}
			if x < w.width-1 {
				sum += area(w, x+1, y, val) // right
			}

			return sum + 1
		}
	}

	return area(w, x, y, val)
}
