package main

import "testing"

func TestCalculateSimpleArea1(t *testing.T) {
	w1 := NewWorld(1, 1)
	a1 := w1.Area(0, 0, EMPTY)
	if a1 != 1 {
		t.Errorf("Area should be 1, got %d", a1)
	}
}

func TestCalculateSimpleArea8(t *testing.T) {
	w2 := NewWorld(2, 4)
	a2 := w2.Area(0, 0, EMPTY)

	if a2 != 8 {
		t.Errorf("Area should be 8, got %d", a2)
	}
}

func TestCalculateSimpleArea10k(t *testing.T) {
	w := NewWorld(100, 100)
	area := w.Area(0, 0, EMPTY)

	if area != 10_000 {
		t.Errorf("Area should be 10000, got %d", area)
	}
}

func TestCalculateSimpleArea100(t *testing.T) {
	w := NewWorld(100, 1)
	area := w.Area(0, 0, EMPTY)

	if area != 100 {
		t.Errorf("Area should be 100, got %d", area)
	}
}

func TestArea1(t *testing.T) {
	// . 1 .
	// . 1 .
	// . 1 .
	w := NewWorld(3, 3)
	w.Set(0, 0, EMPTY)
	w.Set(0, 1, 1)
	w.Set(0, 2, EMPTY)

	w.Set(1, 0, EMPTY)
	w.Set(1, 1, 1)
	w.Set(1, 2, EMPTY)

	w.Set(2, 0, EMPTY)
	w.Set(2, 1, 1)
	w.Set(2, 2, EMPTY)

	area1 := w.Area(0, 0, EMPTY)
	if area1 != 3 {
		t.Errorf("Area should be 3, got %d", area1)
	}

	area2 := w.Area(1, 1, 1)
	if area2 != 3 {
		t.Errorf("Area should be 3, got %d", area2)
	}

	area3 := w.Area(2, 2, EMPTY)
	if area3 != 3 {
		t.Errorf("Area should be 3, got %d", area3)
	}

	area4 := w.Area(1, 1, EMPTY)
	if area4 != 0 {
		t.Errorf("Area should be 0, got %d", area4)
	}
}
