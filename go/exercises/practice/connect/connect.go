package connect

import "errors"

const (
	white = 1 << iota
	black
	connectedWhite
	connectedBlack
)

type colorFlags struct {
	color     int8
	connected int8
}

var flagsBlack = colorFlags{color: black, connected: connectedBlack}
var flagsWhite = colorFlags{color: white, connected: connectedWhite}

type coord struct {
	x, y int
}

type board struct {
	height, width int
	fields        [][]int8
}

func newBoard(lines []string) (board, error) {
	if len(lines) < 1 {
		return board{}, errors.New("no lines given")
	}
	height := len(lines)
	if len(lines[0]) < 1 {
		return board{}, errors.New("first line is empty")
	}
	width := len(lines[0])
	fields := make([][]int8, height)
	backer := make([]int8, height*width)
	for i := range fields {
		fields[i], backer = backer[:width], backer[width:]
	}
	for y, line := range lines {
		for x, c := range line {
			switch c {
			case 'X':
				fields[y][x] = black
			case 'O':
				fields[y][x] = white
			}
		}
	}
	return board{height: height, width: width, fields: fields}, nil
}

func (b board) at(c coord, cf colorFlags) (bool, bool) {
	f := b.fields[c.y][c.x]
	return f&cf.color == cf.color, f&cf.connected == cf.connected
}

func (b board) markConnected(c coord, cf colorFlags) {
	b.fields[c.y][c.x] |= cf.connected
}

func (b board) validCoord(c coord) bool {
	return c.x >= 0 && c.x < b.width && c.y >= 0 && c.y < b.height
}

func (b board) neighbors(c coord) []coord {
	coords := make([]coord, 0, 6)
	dirs := []coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, 1}, {1, -1}}
	for _, d := range dirs {
		nc := coord{x: c.x + d.x, y: c.y + d.y}
		if b.validCoord(nc) {
			coords = append(coords, nc)
		}
	}
	return coords
}

func (b board) startCoords(cf colorFlags) []coord {
	if cf.color == white {
		coords := make([]coord, b.width)
		for i := 0; i < b.width; i++ {
			coords[i] = coord{x: i}
		}
		return coords
	}
	coords := make([]coord, b.height)
	for i := 0; i < b.height; i++ {
		coords[i] = coord{y: i}
	}
	return coords
}

func (b board) isTargetCoord(c coord, cf colorFlags) bool {
	if cf.color == white {
		return c.y == b.height-1
	}
	return c.x == b.width-1
}

func (b board) evaluate(c coord, cf colorFlags) bool {
	stone, connected := b.at(c, cf)
	if stone && !connected {
		b.markConnected(c, cf)
		if b.isTargetCoord(c, cf) {
			return true
		}
		for _, nc := range b.neighbors(c) {
			if b.evaluate(nc, cf) {
				return true
			}
		}
	}
	return false
}

// ResultOf determines the winner of a Hex game board.
func ResultOf(lines []string) (string, error) {
	b, err := newBoard(lines)
	if err != nil {
		return "", err
	}
	for _, c := range b.startCoords(flagsBlack) {
		if b.evaluate(c, flagsBlack) {
			return "X", nil
		}
	}
	for _, c := range b.startCoords(flagsWhite) {
		if b.evaluate(c, flagsWhite) {
			return "O", nil
		}
	}
	return "", nil
}
