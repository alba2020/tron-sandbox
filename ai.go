package main

// -------------- start AI -------------

type AI func(w *World, p *Player) Direction

func AlwaysLeftAI(w *World, p *Player) Direction {
	return Left
}

func RandomDirectionAI(w *World, p *Player) Direction {
	return randomElement(DIRECTIONS)
}

func RandomDirectionNRetriesAI(n int) AI {
	return func(w *World, p *Player) Direction {
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

func RandomValidDirectionAI(w *World, p *Player) Direction {
	valid := w.ValidTurns(p)
	if len(valid) == 0 {
		return randomElement(DIRECTIONS)
	}
	return randomElement(valid)
}

func MaxAreaAI(w *World, p *Player) Direction {
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

func MonteCarloNSimulationsAI(n int) AI {
	return func(world *World, player *Player) Direction {
		decisions := make(map[Direction]int)
		valid := world.ValidTurns(player)

		for _, firstTurn := range valid {
			for i := 0; i < n; i++ {
				simWorld := world.Copy()

				for idx := range simWorld.players {
					simWorld.players[idx].nextTurn = RandomValidDirectionAI
				}

				simWorld.ApplyTurn(player.id, firstTurn) // 1
				simWorld.current = simWorld.NextPlayer(player.id)

				winnerId := simWorld.Simulate()

				if winnerId == player.id {
					decisions[firstTurn]++
				}
			}
		}

		if len(decisions) == 0 { // no ideas :(
			if len(valid) > 0 { // valid turns left
				return randomElement(valid)
			} else {
				return randomElement(DIRECTIONS) // bad turn
			}
		}

		maxDirection, _ := maxValueFromMap(decisions)
		// fmt.Println(decisions)
		// fmt.Printf("MC says %s\n", maxDirection)

		return maxDirection
	}
}

// --------- end AI ------------------
