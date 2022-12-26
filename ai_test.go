package main

import "testing"

func createSmallWorld() (*World, *Player) {
	txt :=
		`. 1 2 . . . . . . .
  	 . 1 2 . . . . . . .
	   . 1 2 . . . . . . .
	   . 1 . . . . . . . .
		 . . . . . . . . . .
		 . . . . . . . . . .
		 . . . . . . . . . .
		 . . . . . . . . . .
		 . . . . . . . . . .
		 . . . . . . . . . .`

	w := ParseWorld(10, 10, txt)
	w.AddPlayer(2, 2, 2, nil)
	player := w.AddPlayer(1, 1, 3, nil)
	return w, player
}

func createBigWorld() (*World, *Player) {
	txt :=
		`. . . . . . . . . . . 1 2 . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . 1 2 . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . 1 2 . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . 1 2 . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . 1 2 . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . 1 2 . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . 1 2 . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . 1 2 . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . 1 . . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . 1 . . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . 1 . . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
		 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .`

	w := ParseWorld(30, 20, txt)
	w.AddPlayer(2, 12, 7, RandomValidDirectionAI)
	player := w.AddPlayer(1, 11, 10, RandomValidDirectionAI)
	w.current = 1
	return w, player
}

func BenchmarkAlwaysLeftAI(b *testing.B) {
	_, p := createSmallWorld()
	p.nextTurn = AlwaysLeftAI

	for i := 0; i < b.N; i++ {
		p.NextTurn()
	}
}

func BenchmarkRandomDirectionAI(b *testing.B) {
	_, p := createSmallWorld()
	p.nextTurn = RandomDirectionAI

	for i := 0; i < b.N; i++ {
		p.NextTurn()
	}
}

func BenchmarkRandomValidDirectionAI(b *testing.B) {
	_, p := createSmallWorld()
	p.nextTurn = RandomValidDirectionAI

	for i := 0; i < b.N; i++ {
		p.NextTurn()
	}
}

func BenchmarkMaxAreaAI(b *testing.B) {
	_, p := createSmallWorld()
	p.nextTurn = MaxAreaAI

	for i := 0; i < b.N; i++ {
		p.NextTurn()
	}
}

func BenchmarkMonteCarloAI(b *testing.B) {
	_, p := createSmallWorld()
	p.nextTurn = MonteCarloNSimulationsAI(5)

	for i := 0; i < b.N; i++ {
		p.NextTurn()
	}
}

// ----- big --------
func BenchmarkMaxAreaAIBig(b *testing.B) {
	_, p := createBigWorld()
	p.nextTurn = MaxAreaAI

	for i := 0; i < b.N; i++ {
		p.NextTurn()
	}
}

func BenchmarkMonteCarloAIBig(b *testing.B) {
	_, p := createBigWorld()
	p.nextTurn = MonteCarloNSimulationsAI(600)

	for i := 0; i < b.N; i++ {
		p.NextTurn()
	}
}
