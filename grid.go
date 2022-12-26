package main

// -------- start grid --------------

type Grid struct {
	width  int8
	height int8
	cells  []int8
}

func (g *Grid) At(x int8, y int8) int8 {
	return g.cells[g.Offset(x, y)]
}

func (g *Grid) Set(x, y int8, val int8) {
	g.cells[g.Offset(x, y)] = val
}

func (g *Grid) Offset(x, y int8) int {
	return int(y)*int(g.width) + int(x)
}

func (g *Grid) Copy() Grid {
	cp := Grid{
		width:  g.width,
		height: g.height,
		cells:  make([]int8, len(g.cells)),
	}
	copy(cp.cells, g.cells)
	return cp
}

func (g *Grid) TurnValid(x, y int8) bool {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		// fmt.Println("out of boundaries")
		return false
	}

	if g.At(x, y) != EMPTY {
		// fmt.Println("not EMPTY")
		return false
	}

	return true
}

func (g *World) ClearPath(playerId int8) {
	for i := range g.cells {
		if g.cells[i] == playerId {
			g.cells[i] = EMPTY
		}
	}
}

func (g *Grid) Area(x, y, val int8) int {
	visited := make(map[int]struct{})

	var area func(g *Grid, x, y, val int8) int

	area = func(g *Grid, x, y, val int8) int {
		if _, found := visited[g.Offset(x, y)]; found {
			return 0
		}
		visited[g.Offset(x, y)] = struct{}{}

		if g.At(x, y) != val {
			return 0
		} else {
			sum := 0
			if y > 0 {
				sum += area(g, x, y-1, val) // up
			}
			if y < g.height-1 {
				sum += area(g, x, y+1, val) // down
			}
			if x > 0 {
				sum += area(g, x-1, y, val) // left
			}
			if x < g.width-1 {
				sum += area(g, x+1, y, val) // right
			}

			return sum + 1
		}
	}

	return area(g, x, y, val)
}

func ParseGrid(w int8, h int8, txt string) Grid {
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

	return Grid{w, h, cells}
}

// --------- end grid -----------------
