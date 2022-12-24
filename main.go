package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func animate(w *World) int8 {
	for w.Active() {
		fmt.Print("\033[H\033[2J") // clear screen
		w.SimulateTurn()
		Show(w)
		time.Sleep(1000 * time.Millisecond)
	}
	winner := w.GetWinner()
	return winner.id
}

func simulate(w, h int8, ais []AI, n int) map[int8]int {
	res := make(map[int8]int)

	for i := 0; i < n; i++ {
		world := NewWorld(w, h)
		for id := range ais {
			world.AddPlayerAtRandom(int8(id), ais[id])
		}
		winId := world.Simulate(0)
		res[winId]++
	}

	return res
}

func doAnimation() {
	w := NewWorld(8, 8)
	w.AddPlayerAtRandom(0, RandomValidDirectionAI)
	w.AddPlayerAtRandom(1, MaxAreaAI)

	winnerId := animate(&w)
	fmt.Printf("game finished, player %d wins\n", winnerId)
}

func doSimulation() {
	ais := []AI{
		RandomValidDirectionAI,
		MaxAreaAI,
	}

	res := simulate(8, 8, ais, 10000)

	var i int8
	for i = 0; i < int8(len(res)); i++ {
		fmt.Printf("player %d wins %d times\n", i, res[i])
	}
}

func main() {
	// doAnimation()
	doSimulation()

	// txt :=
	// 	`. 1 0
	// 	 . 1 .
	// 	 . . .`

	// w := ParseWorld(3, 3, txt)
	// w.Print()
}
