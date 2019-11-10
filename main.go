package main
import (
	"fmt"
	"image/gif"
	"math"
	"os"
	"strings"
)
const (
	border            = true
	topBorderWidth    = 1
	bottomBorderWidth = 1
	leftBorderWidth   = 1
	rightBorderWidth  = 1
)
func f(x float64) (y float64) {
	return math.Sin(x)
}
func main() {
	v := View{
		Graph: Graph{f},
		Min: Point{X: -1, Y: -1.5},
		Max: Point{X: +2, Y: +1.5},
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
		if y < topBorderWidth || y > v.Height-bottomBorderWidth-1 {
			fmt.Print("+" + strings.Repeat("-", v.Width-2) + "+" + "\n")
			continue
		}
		for x := 0; x < v.Width; x++ {
			if x < leftBorderWidth || x > v.Width-rightBorderWidth-1 {
				fmt.Print("|")
				continue
			}
			fmt.Printf("%c", v.AtRune(x, y))
		}
		fmt.Println()
	}
}
