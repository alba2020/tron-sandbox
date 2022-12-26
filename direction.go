package main

// --------- start direction -------

type DirectionMask int8

const (
	None  DirectionMask = 0
	Left  DirectionMask = 1
	Up    DirectionMask = 2
	Right DirectionMask = 4
	Down  DirectionMask = 8
)

func (d DirectionMask) String() string {
	switch d {
	case Left:
		return "LEFT"
	case Up:
		return "UP"
	case Right:
		return "RIGHT"
	case Down:
		return "DOWN"
	}
	return "unknown"
}

var DIRECTIONS = []DirectionMask{
	Left, Up, Right, Down,
}

var DirectionsTable = [16][]DirectionMask{}

func init() {
	DirectionsTable[None] = []DirectionMask{}

	DirectionsTable[Left] = []DirectionMask{Left}
	DirectionsTable[Up] = []DirectionMask{Up}
	DirectionsTable[Right] = []DirectionMask{Right}
	DirectionsTable[Down] = []DirectionMask{Down}

	DirectionsTable[Left|Up] = []DirectionMask{Left, Up}
	DirectionsTable[Up|Right] = []DirectionMask{Up, Right}
	DirectionsTable[Right|Down] = []DirectionMask{Right, Down}
	DirectionsTable[Down|Left] = []DirectionMask{Down, Left}
	DirectionsTable[Left|Right] = []DirectionMask{Left, Right}
	DirectionsTable[Up|Down] = []DirectionMask{Up, Down}

	DirectionsTable[Up|Right|Down] = []DirectionMask{Up, Right, Down}
	DirectionsTable[Left|Right|Down] = []DirectionMask{Left, Right, Down}
	DirectionsTable[Left|Up|Down] = []DirectionMask{Left, Up, Down}
	DirectionsTable[Left|Up|Right] = []DirectionMask{Left, Up, Right}

	DirectionsTable[Left|Up|Right|Down] = []DirectionMask{Left, Up, Right, Down}
}

//  map[DirectionMask][]DirectionMask{
// 	None: {},

// 	Left:  {Left},
// 	Up:    {Up},
// 	Right: {Right},
// 	Down:  {Down},

// 	Left | Up:    {Left, Up},
// 	Up | Right:   {Up, Right},
// 	Right | Down: {Right, Down},
// 	Down | Left:  {Down, Left},
// 	Left | Right: {Left, Right},
// 	Up | Down:    {Up, Down},

// 	Up | Right | Down:   {Up, Right, Down},   // wo Left
// 	Left | Right | Down: {Left, Right, Down}, // wo Up
// 	Left | Up | Down:    {Left, Up, Down},    // wo Right
// 	Left | Up | Right:   {Left, Up, Right},   // wo Down

// 	Left | Up | Right | Down: {Left, Up, Right, Down},
// }

func (d DirectionMask) Random() DirectionMask {
	return randomElement(DirectionsTable[d])
}

// ------ end direction ------------
