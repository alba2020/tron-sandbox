package main

import "testing"

func createSmallWorld() (*World, *Player) {
	txt :=
		`. 1 2 .
  	 . 1 2 .
	   . 1 2 .
	   . 1 . .`

	w := ParseWorld(4, 4, txt)
	player := w.AddPlayer(1, 1, 3, nil)
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
