package main

import "fmt"

var (
	colorReset = "\033[0m"

	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorMagenta = "\033[35m"
	colorCyan    = "\033[36m"
	colorWhite   = "\033[37m"

	colorBrightRed     = "\033[91m"
	colorBrightGreen   = "\033[92m"
	colorBrightYellow  = "\033[93m"
	colorBrightBlue    = "\033[94m"
	colorBrightMagenta = "\033[95m"
	colorBrightCyan    = "\033[96m"
	colorBrightWhite   = "\033[97m"
)

var colors = []string{
	colorRed,
	colorGreen,
	colorYellow,
	colorBlue,
	colorMagenta,
	colorCyan,
	colorWhite,
}

var brightColors = []string{
	colorBrightRed,
	colorBrightGreen,
	colorBrightYellow,
	colorBrightBlue,
	colorBrightMagenta,
	colorBrightCyan,
	colorBrightWhite,
}

func getColor(w *World, x int8, y int8) string {
	id := w.At(x, y)
	for _, p := range w.players {
		if p != nil && x == p.x && y == p.y {
			return brightColors[id]
		}
	}
	return colors[id]
}

func Show(w *World) {
	var i, j int8

	for j = 0; j < w.height; j++ {
		for i = 0; i < w.width; i++ {
			val := w.At(i, j)
			if val == EMPTY {
				fmt.Print(colorReset, ".")
			} else {
				color := getColor(w, i, j)
				fmt.Print(color, val)
			}
		}
		fmt.Println()
	}
	fmt.Print(colorReset)
}

func SimplePrint(w *World) {
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
