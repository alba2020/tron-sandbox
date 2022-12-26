package main

import "testing"

func BenchmarkPointerCopy(b *testing.B) {

	w, _ := createBigWorld()

	for i := 0; i < b.N; i++ {
		cp := w.Copy()
		_ = cp
	}
}

func BenchmarkPlayerAllocation(b *testing.B) {

	const N = 4
	players := [N]*Player{}

	for i := 0; i < b.N; i++ {

		for j := 0; j < N; j++ {
			players[j] = &Player{
				id:       0,
				x:        1,
				y:        2,
				world:    nil,
				nextTurn: AlwaysLeftAI,
			}

		}
	}
}

func BenchmarkPlayerCopy(b *testing.B) {
	const N = 4
	players := [N]Player{}

	for i := 0; i < b.N; i++ {

		for j := 0; j < N; j++ {
			players[j] = Player{
				id:       0,
				x:        1,
				y:        2,
				world:    nil,
				nextTurn: AlwaysLeftAI,
			}
		}

	}
}
