package main

//RuneSet contains a set of runes
type RuneSet struct {
	Empty rune
	Dot   rune

	XAxis   rune
	YAxis   rune
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
	Origin:  '┼',
	HBorder: '─',
	VBorder: '│',

	TopLeftBorder:     '┌',
	TopRightBorder:    '┐',
	BottomLeftBorder:  '└',
	BottomRightBorder: '┘',
}
