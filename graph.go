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
	Gfx   RuneSet

	//Min is the minimum point, it is indicated by the Top-Left Point
	Min Point
	//Max is the maximum point, it is indicated by the Bottom-Right Point
	Max    Point
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
	case v.Gfx.Origin, v.Gfx.YAxis, v.Gfx.XAxis, v.Gfx.Dot:
		return color.Black
	case v.Gfx.Empty:
		return color.White
	}
	panic("Unkown rune")
}
func (v View) AtRune(x, y int) rune {
	xf := +mapValue(x, 0, v.Width, v.Min.X, v.Max.X)
	yf := -mapValue(y, 0, v.Height, v.Min.Y, v.Max.Y)
	xepsilon := (mapValue(1, 0, v.Width, v.Min.X, v.Max.X) - mapValue(0, 0, v.Width, v.Min.X, v.Max.X)) * 0.5
	yepsilon := (mapValue(1, 0, v.Height, v.Min.Y, v.Max.Y) - mapValue(0, 0, v.Height, v.Min.Y, v.Max.X)) * 0.5
	crossed := pass(f, xf, yf, yepsilon)
	if crossed {
		return v.Gfx.Dot
	}
	switch {
	case equal(xf, 0, xepsilon) && equal(yf, 0, yepsilon):
		return v.Gfx.Origin
	case equal(xf, 0, xepsilon):
		return v.Gfx.YAxis
	case equal(yf, 0, yepsilon):
		return v.Gfx.XAxis
	}
	return v.Gfx.Empty
}

func mapValue(v, srcMin, srcMax int, outMin, outMax float64) float64 {
	return outMin + float64(v-srcMin)*(outMax-outMin)/float64(srcMax-srcMin)
}
func equal(a, b float64, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}
func pass(fn func(float64) float64, x, y float64, ydelta float64) bool {
	v := fn(x)
	return math.Abs(y-v) < ydelta
}
