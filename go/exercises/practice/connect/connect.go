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

type coord struct {
	x, y int
}

type board struct {
	height, width int
	fields        [][]int8
}

var (
	flagsBlack = colorFlags{color: black, connected: connectedBlack}
	flagsWhite = colorFlags{color: white, connected: connectedWhite}
)

func newBoard(lines []string) (board, error) {
	if len(lines) == 0 {
		return board{}, errors.New("empty board")
	}
	height := len(lines)
	width := len(lines[0])
	fields := make([][]int8, height)
	for y, line := range lines {
		fields[y] = make([]int8, width)
		for x, ch := range line {
			switch ch {
			case 'X':
				fields[y][x] = black
			case 'O':
				fields[y][x] = white
			}
		}
	}
	return board{height: height, width: width, fields: fields}, nil
}

func (b board) at(c coord, cf colorFlags) (hasColor bool, isConnected bool) {
	v := b.fields[c.y][c.x]
	return v&cf.color != 0, v&cf.connected != 0
}

func (b board) markConnected(c coord, cf colorFlags) {
	b.fields[c.y][c.x] |= cf.connected
}

func (b board) validCoord(c coord) bool {
	return c.x >= 0 && c.x < b.width && c.y >= 0 && c.y < b.height
}

func (b board) neighbors(c coord) []coord {
	dirs := []coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, 1}, {1, -1}}
	result := make([]coord, 0, 6)
	for _, d := range dirs {
		n := coord{c.x + d.x, c.y + d.y}
		if b.validCoord(n) {
			result = append(result, n)
		}
	}
	return result
}

func (b board) startCoords(cf colorFlags) []coord {
	var coords []coord
	if cf.color == white {
		for x := 0; x < b.width; x++ {
			coords = append(coords, coord{x, 0})
		}
	} else {
		for y := 0; y < b.height; y++ {
			coords = append(coords, coord{0, y})
		}
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
	hasColor, isConnected := b.at(c, cf)
	if !hasColor || isConnected {
		return false
	}
	b.markConnected(c, cf)
	if b.isTargetCoord(c, cf) {
		return true
	}
	for _, n := range b.neighbors(c) {
		if b.evaluate(n, cf) {
			return true
		}
	}
	return false
}

// ResultOf determines the winner of a Connect (Hex) game.
// Returns "X" if black wins (left to right), "O" if white wins (top to bottom),
// or "" if there is no winner.
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
