package main

import (
	"fmt"
	"image/gif"
	"math"
	"os"
	"strings"
)

const (
	border = true
)

func f(x float64) (y float64) {
	return math.Sin(x)
}
func main() {
	v := View{
		Graph: Graph{f},
		Gfx:   UnicodeRuneSet,

		Min:    Point{X: -1, Y: -1.5},
		Max:    Point{X: +2, Y: +1.5},
		Width:  40,
		Height: 20,
	}
	f, err := os.Create("graph.gif")
	if err != nil {
		panic(err)
	}
	if err := gif.Encode(f, v, nil); err != nil {
		panic(err)
	}
	for y := 0; y < v.Height; y++ {
		if border && y < 1 {
			fmt.Printf("%c%s%c\n", v.Gfx.TopLeftBorder, strings.Repeat(string(v.Gfx.HBorder), v.Width-2), v.Gfx.TopRightBorder)
			continue
		} else if border && y > v.Height-2 {
			fmt.Printf("%c%s%c\n", v.Gfx.BottomLeftBorder, strings.Repeat(string(v.Gfx.HBorder), v.Width-2), v.Gfx.BottomRightBorder)
			continue
		}
		for x := 0; x < v.Width; x++ {
			if x < 1 || x > v.Width-2 {
				fmt.Print(string(v.Gfx.VBorder))
				continue
			}
			fmt.Printf("%c", v.AtRune(x, y))
		}
		fmt.Println()
	}
}
