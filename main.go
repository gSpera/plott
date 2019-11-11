package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	border = true
)

func main() {
	v := View{
		Min: Point{X: -5, Y: -5},
		Max: Point{X: +5, Y: +5},
	}

	border := flag.Bool("border", true, "enable border, the border will not be added to the width and the height")
	runeset := flag.String("set", "unicode", "runeset to use; valid values: unicode, ascii")
	fn := flag.String("fn", "", "function to plot")
	flag.IntVar(&v.Height, "height", 40, "height of the output graph")
	flag.IntVar(&v.Width, "width", 80, "width of the output graph")
	flag.Var(&v.Min, "min", "top-left of the graph")
	flag.Var(&v.Max, "max", "bottom-right of the graph")
	flag.Parse()

	switch strings.ToLower(*runeset) {
	case "unicode":
		v.Gfx = UnicodeRuneSet
	case "ascii":
		v.Gfx = ASCIIRuneSet
	default:
		fmt.Fprintf(os.Stderr, "Unsupported runeset: %q", *runeset)
		os.Exit(1)
	}

	if *fn == "" {
		fmt.Println("Function not supplied")
		flag.PrintDefaults()
		os.Exit(1)
	}

	v.Graph.fn = FunctionEvaluator(*fn)

	for y := 0; y < v.Height; y++ {
		if *border && y < 1 {
			fmt.Printf("%c%s%c\n", v.Gfx.TopLeftBorder, strings.Repeat(string(v.Gfx.HBorder), v.Width-2), v.Gfx.TopRightBorder)
			continue
		} else if *border && y > v.Height-2 {
			fmt.Printf("%c%s%c\n", v.Gfx.BottomLeftBorder, strings.Repeat(string(v.Gfx.HBorder), v.Width-2), v.Gfx.BottomRightBorder)
			continue
		}
		for x := 0; x < v.Width; x++ {
			if *border && (x < 1 || x > v.Width-2) {
				fmt.Print(string(v.Gfx.VBorder))
				continue
			}
			fmt.Printf("%c", v.AtRune(x, y))
		}
		fmt.Println()
	}
}
