package counter

import "testing"

func assertCounts(t *testing.T, c Counter, lines, letters, chars int) {
	t.Helper()
	if got := c.Lines(); got != lines {
		t.Errorf("Lines: got %d, want %d", got, lines)
	}
	if got := c.Letters(); got != letters {
		t.Errorf("Letters: got %d, want %d", got, letters)
	}
	if got := c.Characters(); got != chars {
		t.Errorf("Characters: got %d, want %d", got, chars)
	}
}

func TestNoAddString(t *testing.T) {
	c := makeCounter()
	assertCounts(t, c, 0, 0, 0)
}

func TestEmptyString(t *testing.T) {
	c := makeCounter()
	c.AddString("")
	assertCounts(t, c, 0, 0, 0)
}

func TestSimpleASCIINoNewline(t *testing.T) {
	c := makeCounter()
	c.AddString("hello")
	assertCounts(t, c, 1, 5, 5)
}

func TestASCIIWithNewlineInMiddle(t *testing.T) {
	c := makeCounter()
	c.AddString("Hello\nworld!")
	assertCounts(t, c, 2, 10, 12)
}

func TestStringEndingWithNewline(t *testing.T) {
	c := makeCounter()
	c.AddString("hello\n")
	assertCounts(t, c, 1, 5, 6)
}

func TestUnicodeLetters(t *testing.T) {
	c := makeCounter()
	c.AddString("здравствуй, мир\n")
	assertCounts(t, c, 1, 13, 16)
}

func TestMultipleAddStrings(t *testing.T) {
	c := makeCounter()
	c.AddString("hello\n")
	c.AddString("world")
	assertCounts(t, c, 2, 10, 11)
}

func TestOnlyNewlines(t *testing.T) {
	c := makeCounter()
	c.AddString("\n\n\n")
	assertCounts(t, c, 3, 0, 3)
}

func TestMixedContent(t *testing.T) {
	c := makeCounter()
	c.AddString("abc 123!@#\ndef")
	assertCounts(t, c, 2, 6, 14)
}
