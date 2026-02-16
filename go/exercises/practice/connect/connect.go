package connect

import "errors"

type point struct {
	row, col int
}

func ResultOf(lines []string) (string, error) {
	if len(lines) == 0 {
		return "", errors.New("empty board")
	}
	if len(lines[0]) == 0 {
		return "", errors.New("empty first line")
	}
	if hasWon(lines, 'X') {
		return "X", nil
	}
	if hasWon(lines, 'O') {
		return "O", nil
	}
	return "", nil
}

func hasWon(lines []string, player byte) bool {
	rows := len(lines)
	cols := len(lines[0])
	visited := make(map[point]bool)
	queue := []point{}

	if player == 'X' {
		for r := 0; r < rows; r++ {
			if lines[r][0] == player {
				p := point{r, 0}
				queue = append(queue, p)
				visited[p] = true
			}
		}
	} else {
		for c := 0; c < cols; c++ {
			if lines[0][c] == player {
				p := point{0, c}
				queue = append(queue, p)
				visited[p] = true
			}
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if player == 'X' && cur.col == cols-1 {
			return true
		}
		if player == 'O' && cur.row == rows-1 {
			return true
		}
		for _, nb := range neighbors(cur, rows, cols) {
			if !visited[nb] && lines[nb.row][nb.col] == player {
				visited[nb] = true
				queue = append(queue, nb)
			}
		}
	}
	return false
}

func neighbors(p point, rows, cols int) []point {
	deltas := [6]point{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, -1}, {-1, 1},
	}
	result := make([]point, 0, 6)
	for _, d := range deltas {
		nr, nc := p.row+d.row, p.col+d.col
		if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
			result = append(result, point{nr, nc})
		}
	}
	return result
}
