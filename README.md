# Plot To Text
Program Plot Text is used to plot math graphs in a textual rapresentation with monospaced fonts.

## Function

PlotT operates with command-line flag, the most important is -fn,
this defines the function to plot, the function is dependent on x,
see github.com/gSpera/meval for the functions and the constants that can be used
```
plott -fn="sin(x)"
┌────────────────────────────┐
│              │             │
│              │             │
│              │             │
│              │             │
│              │             │
│              │             │
│              │             │
│···           │  ····       │
│   ·          │··    ··     │
│────··────────·────────··───│
│      ··    ··│          ·  │
│        ····  │           ··│
│              │             │
│              │             │
│              │             │
│              │             │
│              │             │
│              │             │
└────────────────────────────┘
```

## Size

By defaults plott outputs to a 80-column format(80x40), this can be changed width the -width and -height flags.
The plot is based around the Origin(0; 0) and spans from (-5; 5)(top-left) to (5; -5)(bottom-right),
this behaviour can be changed with the -min="(X; Y)"(top-left) and the -max="(X; Y)"(bottom-right) flags.
The examples in this document are created(if not otherwise specified) in 30x20 runes.
```
plott -fn="sin(x)" -min="(-3.14; 1.5)" -max="(3.14; -1.5)"
┌────────────────────────────┐
│              │             │
│              │             │
│              │      ··     │
│              │    ··  ··   │
│              │   ·      ·  │
│              │  ·        · │
│              │ ·          ·│
│              │             │
│              │·            │
│──────────────·─────────────│
│·            ·│             │
│              │             │
│ ·          · │             │
│  ·        ·  │             │
│   ·      ·   │             │
│    ··  ··    │             │
│      ··      │             │
│              │             │
└────────────────────────────┘
```

## Unicode

PlotT uses Unicode runes by default, if the output device doesn't support Unicode the -set=ascii flag may be used.
other sets(and custom ones) may be avaible in future.
```
plott -fn="x*x" -set=ascii
+----------------------------+
|..............|.............|
|..............|.............|
|........x.....|.....x.......|
|..............|.............|
|..............|.............|
|.........x....|....x........|
|..............|.............|
|..............|.............|
|..........x...|...x.........|
|..............|.............|
|...........x..|..x..........|
|..............|.............|
|............x.|.x...........|
|-------------xxx------------|
|..............|.............|
|..............|.............|
|..............|.............|
|..............|.............|
+----------------------------+
```

## Border

The plot is surrounded by a border by default, this can be changed with the -border=false flag
```
plott -fn="tan(x)" -set=ascii -border=false
...............|..............
...............|..............
...............|...x..........
x..............|..............
...............|..............
...............|............x.
.........x.....|..............
...............|..x...........
........x......|.x.........x..
.......x.......|x.........x...
-----xx--------x--------xx----
....x.........x|.......x......
...x.........x.|......x.......
............x..|..............
...............|.....x........
..x............|..............
...............|..............
...............|..............
...........x...|..............
...............|..............
```
