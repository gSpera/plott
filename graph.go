package main

import (
	"image"
	"image/color"
	"math"
)

type Point struct {
	X float64
	Y float64
}

type Function func(x float64) (y float64)

//Graph rapresent a math graph
type Graph struct {
	fn Function
}

//View is a view of a graph
type View struct {
	Graph Graph

	//Min is the minimum point, it is indicated by the Top-Left Point
	Min Point
	//Max is the maximum point, it is indicated by the Bottom-Right Point
	Max Point

	Width  int
	Height int
}

func (v View) Bounds() image.Rectangle {
	return image.Rect(0, 0, v.Width, v.Height)
}

func (v View) ColorModel() color.Model {
	return color.RGBAModel
}

func (v View) At(x, y int) color.Color {
	switch v.AtRune(x, y) {
	case '+', '|', '-', 'x':
		return color.Black
	case '.':
		return color.White
	}
	panic("Unkown rune")
}

func (v View) AtRune(x, y int) rune {
	xf := +mapValue(x, 0, v.Width, v.Min.X, v.Max.X)
	yf := -mapValue(y, 0, v.Height, v.Min.Y, v.Max.Y)
	xepsilon := (v.Max.X - v.Min.X) / float64(v.Width)
	yepsilon := (v.Max.Y - v.Min.Y) / float64(v.Height)
	crossed := pass(f, xf, yf, yepsilon)

	if crossed {
		return 'x'
	}
	switch {
	case equal(xf, 0, xepsilon*xEpsilonScale) && equal(yf, 0, yepsilon*yEpsilonScale):
		return '+'
	case equal(xf, 0, xepsilon*xEpsilonScale):
		return '|'
	case equal(yf, 0, yepsilon*yEpsilonScale):
		return '-'
	}
	return '.'
}

func mapValue(v, srcMin, srcMax int, outMin, outMax float64) float64 {
	return outMin + float64(v-srcMin)*(outMax-outMin)/float64(srcMax-srcMin)
}

func equal(a, b float64, epsilon float64) bool {
	return math.Abs(a-b) < epsilon*epsilonScale
}
func pass(fn func(float64) float64, x, y float64, epsilon float64) bool {
	v := fn(x)
	return math.Abs(y-v) < epsilon*epsilonScale
}
