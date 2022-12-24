package main

type AI func(w *World, p *Player) string

var DIRECTIONS = []string{
	"UP", "DOWN", "LEFT", "RIGHT",
}

func AlwaysLeftAI(w *World, p *Player) string {
	return "LEFT"
}

func RandomDirectionAI(w *World, p *Player) string {
	return randomElement(DIRECTIONS)
}

func RandomDirectionNRetriesAI(n int) AI {
	return func(w *World, p *Player) string {
		for i := 0; i < n; i++ {
			turn := randomElement(DIRECTIONS)
			x, y := p.NextCoords(turn)
			if w.TurnValid(x, y) {
				return turn
			}
		}
		return randomElement(DIRECTIONS)
	}
}

func RandomValidDirectionAI(w *World, p *Player) string {
	valid := w.ValidTurns(p)
	if len(valid) == 0 {
		return randomElement(DIRECTIONS)
	}
	return randomElement(valid)
}

func MaxAreaAI(w *World, p *Player) string {
	valid := w.ValidTurns(p)
	if len(valid) == 0 {
		return randomElement(DIRECTIONS)
	}
	turn := randomElement(valid)
	maxArea := 0

	for _, dir := range valid {
		next_x, next_y := p.NextCoords(dir)
		area := w.Area(next_x, next_y, EMPTY)
		if area > maxArea {
			maxArea = area
			turn = dir
		}
	}

	return turn
}

func MonteCarloMaxAreaAI(w *World, p *Player) string {
	valid := w.ValidTurns(p)
	for _, validTurn := range valid {
		// n simulations
		simWorld := w.Copy()
		simPlayer := w.GetPlayer(p.id)
		simWorld.ApplyTurn(simPlayer, validTurn)
	}

	return ""
}
