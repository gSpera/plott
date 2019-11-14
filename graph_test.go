package main

import (
	"fmt"
	"image"
	"image/color"
	"testing"
)

func TestPoint_Set(t *testing.T) {
	tm := []struct {
		name   string
		input  string
		output Point
	}{
		{
			"canonical",
			"(-5; +5)",
			Point{X: -5, Y: 5},
		},
		{
			"no space",
			"(7;-7)",
			Point{X: 7, Y: -7},
		},
		{
			"many spaces",
			"(   8;     +7)",
			Point{X: 8, Y: 7},
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{}
			err := p.Set(tt.input)

			if err != nil {
				t.Errorf("Error while decoding Point: %v", err)
			}
			if p != tt.output {
				t.Errorf("Points differs for input %q: expected: %+v; got: %+v", tt.input, tt.output, p)
			}
		})
	}
}

func TestPoint_String(t *testing.T) {
	tm := []struct {
		name  string
		out   string
		point Point
	}{
		{
			name:  "1;1",
			out:   "(1.00; 1.00)",
			point: Point{X: 1, Y: 1},
		},
		{
			name:  "origin",
			out:   "(0.00; 0.00)",
			point: Point{X: 0, Y: 0},
		},
		{
			name:  "negative",
			out:   "(-5.00; 7.00)",
			point: Point{X: -5, Y: 7},
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.point.String()

			if got != tt.out {
				t.Errorf("Wrong Value: expected: %q; got: %q", tt.out, got)
			}
		})
	}
}

func TestView_image(t *testing.T) {
	t.Run("bounds", func(t *testing.T) {
		v := View{
			Width:  1,
			Height: 2,
		}
		expect := image.Rect(0, 0, 1, 2)
		got := v.Bounds()

		if got != expect {
			t.Errorf("Wrong Rect: expected: %+v; got: %+v", expect, got)
		}
	})

	t.Run("color model", func(t *testing.T) {
		v := View{}
		expect := color.RGBAModel
		got := v.ColorModel()

		if got != expect {
			t.Errorf("Wrong Color Model: expected: %+v; got: %+v", expect, got)
		}
	})
}

func TestViewAt(t *testing.T) {
	f, _ := FunctionEvaluator("x+1")
	v := View{
		Function: f,
		Width:    10,
		Height:   10,
		Border:   true,
		Axis:     true,
		Min:      Point{X: -5, Y: 5},
		Max:      Point{X: 5, Y: -5},
		Gfx:      ASCIIRuneSet,
	}

	tm := []struct {
		name        string
		x           int
		y           int
		expectRune  rune
		expectColor color.Color
	}{
		{
			"top-left",
			1, 1,
			v.Gfx.Empty,
			color.White,
		},
		{
			"origin",
			5, 5,
			v.Gfx.Origin,
			color.Black,
		},
		{
			"x-axis",
			5, 0,
			v.Gfx.XAxis,
			color.Black,
		},
		{
			"y-axis",
			0, 5,
			v.Gfx.YAxis,
			color.Black,
		},
		{
			"dot",
			5, 4,
			v.Gfx.Dot,
			color.Black,
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			gotRune := v.AtRune(tt.x, tt.y)
			if gotRune != tt.expectRune {
				t.Errorf("Expected Rune: %c(%x); got: %c(%x)", tt.expectRune, tt.expectRune, gotRune, gotRune)
			}

			gotColor := v.At(tt.x, tt.y)
			if gotColor != tt.expectColor {
				t.Errorf("Expected Color: %+v; got: %+v", tt.expectColor, gotColor)
			}
		})
	}
}

func TestFunctionEvaluator(t *testing.T) {
	f, err := FunctionEvaluator("x")
	if err != nil {
		t.Errorf("cannot evaluate: %w", err)
		return
	}

	if got := f(1); got != 1 {
		t.Errorf("f(x) = x; f(1) != 1; got: %f", got)
	}

	_, err = FunctionEvaluator("y")
	if err == nil {
		t.Errorf("Evaluating wrong function, expected error; got nil")
	}
}

func TestMapValue(t *testing.T) {
	tm := []struct {
		name string
		out  float64

		v      int
		minSrc int
		maxSrc int
		minOut float64
		maxOut float64
	}{
		{
			name: "min",
			out:  1,

			v:      10,
			minSrc: 10,
			maxSrc: 100,
			minOut: 1,
			maxOut: 100,
		},
		{
			name: "max",
			out:  200,

			v:      10,
			minSrc: 1,
			maxSrc: 10,
			minOut: 2,
			maxOut: 200,
		},
		{
			name: "middle",
			out:  5,

			v:      50,
			minSrc: 10,
			maxSrc: 100,
			minOut: 1,
			maxOut: 10,
		},
		{
			name: "over",
			out:  11,

			v:      110,
			minSrc: 10,
			maxSrc: 100,
			minOut: 1,
			maxOut: 10,
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			got := mapValue(tt.v, tt.minSrc, tt.maxSrc, tt.minOut, tt.maxOut)

			if got != tt.out {
				t.Errorf("Wrong Value: expected: %f; got: %f", tt.out, got)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	tm := []struct {
		a, b    float64
		epsilon float64
		expect  bool
	}{
		{
			1,
			1.5,
			0.9,
			true,
		},
		{
			1,
			1.1,
			0.0001,
			false,
		},
	}

	for _, tt := range tm {
		t.Run(fmt.Sprintf("%f == %f (%f epsilon)", tt.a, tt.b, tt.epsilon), func(t *testing.T) {
			got := equal(tt.a, tt.b, tt.epsilon)
			if got != tt.expect {
				t.Errorf("Wrong Value: %f == %f (%f epsilon) expected: %t; got: %t", tt.a, tt.b, tt.epsilon, tt.expect, got)
			}
		})
	}
}

func TestPass(t *testing.T) {
	fx := func(x float64) float64 { return x }
	tm := []struct {
		f      func(float64) float64
		x, y   float64
		delta  float64
		expect bool
	}{
		{
			fx,
			1,
			1,
			0.01,
			true,
		},
		{
			fx,
			1,
			2,
			0.1,
			false,
		},
	}

	for _, tt := range tm {
		t.Run(fmt.Sprintf("f(x) pass (%f; %f) (%f epsilon)", tt.x, tt.y, tt.delta), func(t *testing.T) {
			got := pass(tt.f, tt.x, tt.y, tt.delta)
			if got != tt.expect {
				t.Errorf("expected: %t; got: %t", tt.expect, got)
			}
		})
	}
}
