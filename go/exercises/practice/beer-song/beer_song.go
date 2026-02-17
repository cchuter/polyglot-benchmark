package beer

import (
	"fmt"
	"strings"
)

func Verse(n int) (string, error) {
	if n < 0 || n > 99 {
		return "", fmt.Errorf("invalid verse number: %d", n)
	}
	switch n {
	case 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	case 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	case 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	default:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
	}
}

func Verses(upper, lower int) (string, error) {
	if upper > 99 || lower < 0 || upper < lower {
		return "", fmt.Errorf("invalid range: %d to %d", upper, lower)
	}
	var parts []string
	for i := upper; i >= lower; i-- {
		v, _ := Verse(i)
		parts = append(parts, v)
	}
	return strings.Join(parts, "\n") + "\n", nil
}

func Song() string {
	s, _ := Verses(99, 0)
	return s
}
