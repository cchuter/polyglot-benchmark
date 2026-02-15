package bowling

import "errors"

// Game tracks the state of a bowling game.
type Game struct {
	rolls       []int
	frame       int
	rollInFrame int
	done        bool
}

// NewGame returns a new Game instance.
func NewGame() *Game {
	return &Game{}
}

// Roll records a roll of the given number of pins.
func (g *Game) Roll(pins int) error {
	if g.done {
		return errors.New("Cannot roll after game is over")
	}
	if pins < 0 {
		return errors.New("Negative roll is invalid")
	}
	if pins > 10 {
		return errors.New("Pin count exceeds pins on the lane")
	}

	if g.frame < 9 {
		// Frames 0-8: second roll in frame must not exceed remaining pins
		if g.rollInFrame == 1 {
			if g.rolls[len(g.rolls)-1]+pins > 10 {
				return errors.New("Pin count exceeds pins on the lane")
			}
		}
	} else {
		// Frame 9 (10th frame)
		if g.rollInFrame == 1 {
			prev := g.rolls[len(g.rolls)-1]
			if prev != 10 && prev+pins > 10 {
				return errors.New("Pin count exceeds pins on the lane")
			}
		} else if g.rollInFrame == 2 {
			first := g.rolls[len(g.rolls)-2]
			second := g.rolls[len(g.rolls)-1]
			if first == 10 && second == 10 {
				// Both strikes, pins 0-10 is fine
			} else if first == 10 && second != 10 {
				if second+pins > 10 {
					return errors.New("Pin count exceeds pins on the lane")
				}
			}
			// Spare (first+second==10, first!=10): pins 0-10 is fine
		}
	}

	g.rolls = append(g.rolls, pins)

	// Advance frame state
	if g.frame < 9 {
		if pins == 10 && g.rollInFrame == 0 {
			g.frame++
			g.rollInFrame = 0
		} else if g.rollInFrame == 1 {
			g.frame++
			g.rollInFrame = 0
		} else {
			g.rollInFrame = 1
		}
	} else {
		if g.rollInFrame == 0 {
			g.rollInFrame = 1
		} else if g.rollInFrame == 1 {
			first := g.rolls[len(g.rolls)-2]
			if first == 10 || first+pins == 10 {
				g.rollInFrame = 2
			} else {
				g.done = true
			}
		} else {
			g.done = true
		}
	}

	return nil
}

// Score returns the total score if the game is complete.
func (g *Game) Score() (int, error) {
	if !g.done {
		return 0, errors.New("Score cannot be taken until the end of the game")
	}

	score := 0
	i := 0
	for frame := 0; frame < 9; frame++ {
		if g.rolls[i] == 10 {
			score += 10 + g.rolls[i+1] + g.rolls[i+2]
			i++
		} else if g.rolls[i]+g.rolls[i+1] == 10 {
			score += 10 + g.rolls[i+2]
			i += 2
		} else {
			score += g.rolls[i] + g.rolls[i+1]
			i += 2
		}
	}

	// 10th frame: add remaining rolls
	for ; i < len(g.rolls); i++ {
		score += g.rolls[i]
	}

	return score, nil
}
