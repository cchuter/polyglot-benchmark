package connect

type coord struct {
	row, col int
}

type board struct {
	height int
	width  int
	cells  []string
}

func newBoard(lines []string) board {
	h := len(lines)
	w := 0
	if h > 0 {
		w = len(lines[0])
	}
	return board{height: h, width: w, cells: lines}
}

func (b *board) at(c coord) byte {
	return b.cells[c.row][c.col]
}

func (b *board) valid(c coord) bool {
	return c.row >= 0 && c.row < b.height && c.col >= 0 && c.col < b.width
}

func (b *board) neighbors(c coord) []coord {
	dirs := [6]coord{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {-1, 1}, {1, -1},
	}
	result := make([]coord, 0, 6)
	for _, d := range dirs {
		n := coord{c.row + d.row, c.col + d.col}
		if b.valid(n) {
			result = append(result, n)
		}
	}
	return result
}

func (b *board) dfs(c coord, color byte, visited map[coord]bool) bool {
	if !b.valid(c) || visited[c] || b.at(c) != color {
		return false
	}
	visited[c] = true
	if color == 'X' && c.col == b.width-1 {
		return true
	}
	if color == 'O' && c.row == b.height-1 {
		return true
	}
	for _, n := range b.neighbors(c) {
		if b.dfs(n, color, visited) {
			return true
		}
	}
	return false
}

// ResultOf determines the winner of a Hex board game.
// Returns "X" if X connects left to right, "O" if O connects top to bottom, or "" if no winner.
func ResultOf(lines []string) (string, error) {
	b := newBoard(lines)
	if b.height == 0 || b.width == 0 {
		return "", nil
	}

	// Check if X wins (left to right)
	visited := make(map[coord]bool)
	for row := 0; row < b.height; row++ {
		c := coord{row, 0}
		if b.at(c) == 'X' && b.dfs(c, 'X', visited) {
			return "X", nil
		}
	}

	// Check if O wins (top to bottom)
	visited = make(map[coord]bool)
	for col := 0; col < b.width; col++ {
		c := coord{0, col}
		if b.at(c) == 'O' && b.dfs(c, 'O', visited) {
			return "O", nil
		}
	}

	return "", nil
}
