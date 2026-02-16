package beer

import (
	"bytes"
	"fmt"
)

// Song returns the full lyrics for 99 bottles of beer.
func Song() string {
	result, _ := Verses(99, 0)
	return result
}

// Verses returns the lyrics from verse start down to verse stop.
func Verses(start, stop int) (string, error) {
	if start < 0 || start > 99 {
		return "", fmt.Errorf("start value[%d] is not a valid verse", start)
	}
	if stop < 0 || stop > 99 {
		return "", fmt.Errorf("stop value[%d] is not a valid verse", stop)
	}
	if start < stop {
		return "", fmt.Errorf("start value[%d] is less than stop value[%d]", start, stop)
	}

	var buf bytes.Buffer
	for i := start; i >= stop; i-- {
		v, _ := Verse(i)
		buf.WriteString(v)
		buf.WriteString("\n")
	}
	return buf.String(), nil
}

// Verse returns a single verse of the song.
func Verse(n int) (string, error) {
	switch {
	case n < 0 || n > 99:
		return "", fmt.Errorf("%d is not a valid verse", n)
	case n == 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	case n == 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	case n == 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	default:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
	}
}
