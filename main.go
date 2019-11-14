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
		Min: Point{X: -5, Y: +5},
		Max: Point{X: +5, Y: -5},
	}

	border := flag.Bool("border", true, "enable border, the border will not be added to the width and the height")
	axis := flag.Bool("axis", true, "enable x and y axis")
	runeset := flag.String("set", "unicode", "runeset to use; valid values: unicode, ascii")
	fn := flag.String("fn", "", "function to plot")
	flag.IntVar(&v.Height, "height", 40, "height of the output graph")
	flag.IntVar(&v.Width, "width", 80, "width of the output graph")
	flag.Var(&v.Min, "min", "top-left of the graph")
	flag.Var(&v.Max, "max", "bottom-right of the graph")
	flag.Parse()

	v.Border = *border
	v.Axis = *axis

	switch strings.ToLower(*runeset) {
	case "unicode":
		v.Gfx = UnicodeRuneSet
	case "ascii":
		v.Gfx = ASCIIRuneSet
	default:
		fmt.Fprintf(os.Stderr, "Unsupported runeset: %q\n", *runeset)
		os.Exit(1)
	}

	if *fn == "" {
		fmt.Println("Function not supplied")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var err error
	v.Function, err = FunctionEvaluator(*fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot evaluate function %q: %v", *fn, err)
		os.Exit(1)
	}

	for y := 0; y < v.Height; y++ {
		for x := 0; x < v.Width; x++ {
			fmt.Printf("%c", v.AtRune(x, y))
		}
		fmt.Println()
	}
}
