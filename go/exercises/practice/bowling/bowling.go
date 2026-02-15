package bowling

import "errors"

// Game tracks the state of a bowling game.
type Game struct {
	rolls        []int
	currentFrame int
	currentBall  int
	done         bool
}

// NewGame creates a new bowling game.
func NewGame() *Game {
	return &Game{}
}

// Roll records a roll of the ball knocking down the given number of pins.
func (g *Game) Roll(pins int) error {
	if g.done {
		return errors.New("cannot roll after game is over")
	}
	if pins < 0 {
		return errors.New("negative roll is invalid")
	}
	if pins > 10 {
		return errors.New("pin count exceeds pins on the lane")
	}

	if g.currentFrame < 9 {
		if g.currentBall == 0 {
			if pins == 10 {
				g.rolls = append(g.rolls, pins)
				g.currentFrame++
			} else {
				g.rolls = append(g.rolls, pins)
				g.currentBall = 1
			}
		} else {
			prev := g.rolls[len(g.rolls)-1]
			if prev+pins > 10 {
				return errors.New("pin count exceeds pins on the lane")
			}
			g.rolls = append(g.rolls, pins)
			g.currentFrame++
			g.currentBall = 0
		}
	} else {
		switch g.currentBall {
		case 0:
			g.rolls = append(g.rolls, pins)
			g.currentBall = 1
		case 1:
			firstRoll := g.rolls[len(g.rolls)-1]
			if firstRoll == 10 {
				// First was a strike, pins reset
			} else if firstRoll+pins > 10 {
				return errors.New("pin count exceeds pins on the lane")
			}
			g.rolls = append(g.rolls, pins)
			if firstRoll == 10 || firstRoll+pins == 10 {
				g.currentBall = 2
			} else {
				g.done = true
			}
		case 2:
			firstRoll := g.rolls[len(g.rolls)-2]
			secondRoll := g.rolls[len(g.rolls)-1]
			if firstRoll == 10 && secondRoll == 10 {
				// Two strikes, pins reset
			} else if firstRoll == 10 {
				// First was strike, second was not; pins not reset
				if secondRoll+pins > 10 {
					return errors.New("pin count exceeds pins on the lane")
				}
			}
			// Spare case: pins reset, 0-10 valid (already checked by top-level)
			g.rolls = append(g.rolls, pins)
			g.done = true
		}
	}

	return nil
}

// Score returns the total score for a completed game.
func (g *Game) Score() (int, error) {
	if !g.done {
		return 0, errors.New("score cannot be taken until the end of the game")
	}

	score := 0
	rollIndex := 0

	for frame := 0; frame < 10; frame++ {
		if g.rolls[rollIndex] == 10 {
			score += 10 + g.rolls[rollIndex+1] + g.rolls[rollIndex+2]
			rollIndex++
		} else if g.rolls[rollIndex]+g.rolls[rollIndex+1] == 10 {
			score += 10 + g.rolls[rollIndex+2]
			rollIndex += 2
		} else {
			score += g.rolls[rollIndex] + g.rolls[rollIndex+1]
			rollIndex += 2
		}
	}

	return score, nil
}
