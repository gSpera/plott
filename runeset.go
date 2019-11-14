package main

//RuneSet contains a set of runes
type RuneSet struct {
	Empty rune
	Dot   rune

	XAxis   rune
	YAxis   rune
	XArrow  rune
	YArrow  rune
	Origin  rune
	HBorder rune
	VBorder rune

	TopLeftBorder     rune
	TopRightBorder    rune
	BottomLeftBorder  rune
	BottomRightBorder rune
}

var ASCIIRuneSet = RuneSet{
	Empty: '.',
	Dot:   'x',

	XAxis:   '-',
	YAxis:   '|',
	XArrow:  '>',
	YArrow:  '^',
	Origin:  '+',
	HBorder: '-',
	VBorder: '|',

	TopLeftBorder:     '+',
	TopRightBorder:    '+',
	BottomLeftBorder:  '+',
	BottomRightBorder: '+',
}

var UnicodeRuneSet = RuneSet{
	Empty: ' ',
	Dot:   '·',

	XAxis:   '─',
	YAxis:   '│',
	XArrow:  '►',
	YArrow:  '▲',
	Origin:  '┼',
	HBorder: '─',
	VBorder: '│',

	TopLeftBorder:     '┌',
	TopRightBorder:    '┐',
	BottomLeftBorder:  '└',
	BottomRightBorder: '┘',
}
