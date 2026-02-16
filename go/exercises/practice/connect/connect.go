package connect

import "errors"

type coord struct {
	x, y int
}

type board struct {
	height int
	width  int
	fields [][]byte
}

func newBoard(lines []string) (board, error) {
	if len(lines) == 0 {
		return board{}, errors.New("no lines given")
	}
	if len(lines[0]) == 0 {
		return board{}, errors.New("first line is empty")
	}
	height := len(lines)
	width := len(lines[0])
	fields := make([][]byte, height)
	for y, line := range lines {
		fields[y] = []byte(line)
	}
	return board{height: height, width: width, fields: fields}, nil
}

func (b board) validCoord(c coord) bool {
	return c.x >= 0 && c.x < b.width && c.y >= 0 && c.y < b.height
}

func (b board) neighbors(c coord) []coord {
	dirs := []coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, 1}, {1, -1}}
	result := make([]coord, 0, 6)
	for _, d := range dirs {
		nc := coord{x: c.x + d.x, y: c.y + d.y}
		if b.validCoord(nc) {
			result = append(result, nc)
		}
	}
	return result
}

func (b board) hasPath(start coord, stone byte, isTarget func(coord) bool, visited map[coord]bool) bool {
	if visited[start] {
		return false
	}
	if b.fields[start.y][start.x] != stone {
		return false
	}
	visited[start] = true
	if isTarget(start) {
		return true
	}
	for _, nc := range b.neighbors(start) {
		if b.hasPath(nc, stone, isTarget, visited) {
			return true
		}
	}
	return false
}

// ResultOf determines the winner of a Hex board game.
// "X" wins by connecting left to right, "O" wins by connecting top to bottom.
func ResultOf(lines []string) (string, error) {
	b, err := newBoard(lines)
	if err != nil {
		return "", err
	}
	// Check if X wins (left to right)
	visited := make(map[coord]bool)
	for y := 0; y < b.height; y++ {
		start := coord{x: 0, y: y}
		if b.hasPath(start, 'X', func(c coord) bool { return c.x == b.width-1 }, visited) {
			return "X", nil
		}
	}
	// Check if O wins (top to bottom)
	visited = make(map[coord]bool)
	for x := 0; x < b.width; x++ {
		start := coord{x: x, y: 0}
		if b.hasPath(start, 'O', func(c coord) bool { return c.y == b.height-1 }, visited) {
			return "O", nil
		}
	}
	return "", nil
}
