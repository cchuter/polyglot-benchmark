// Package bowling implements scoring for the game of bowling.
package bowling

import "errors"

var (
	ErrNegativeRollIsInvalid        = errors.New("Negative roll is invalid")
	ErrPinCountExceedsPinsOnTheLane = errors.New("Pin count exceeds pins on the lane")
	ErrPrematureScore               = errors.New("Score cannot be taken until the end of the game")
	ErrCannotRollAfterGameOver      = errors.New("Cannot roll after game is over")
)

const (
	pinsPerFrame      = 10
	framesPerGame     = 10
	maxRollsPerFrame  = 2
	maxRollsLastFrame = 3
	maxRolls          = (maxRollsPerFrame * (framesPerGame - 1)) + maxRollsLastFrame
)

type Game struct {
	rolls       [maxRolls]int
	nRolls      int
	nFrames     int
	rFrameStart int
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Roll(pins int) error {
	if pins > pinsPerFrame {
		return ErrPinCountExceedsPinsOnTheLane
	}
	if pins < 0 {
		return ErrNegativeRollIsInvalid
	}
	if g.completedFrames() == framesPerGame {
		return ErrCannotRollAfterGameOver
	}
	g.rolls[g.nRolls] = pins
	g.nRolls++
	if pins == pinsPerFrame && g.completedFrames() < framesPerGame-1 {
		g.completeTheFrame()
		return nil
	}
	if g.rollsThisFrame() == maxRollsPerFrame {
		if g.rawFrameScore(g.rFrameStart) > pinsPerFrame {
			if g.completedFrames() != framesPerGame-1 || !g.isStrike(g.rFrameStart) {
				return ErrPinCountExceedsPinsOnTheLane
			}
		}
		if g.completedFrames() < framesPerGame-1 {
			g.completeTheFrame()
			return nil
		}
		if g.rawFrameScore(g.rFrameStart) < pinsPerFrame {
			g.completeTheFrame()
		}
	} else if g.rollsThisFrame() == maxRollsLastFrame {
		if g.isStrike(g.rFrameStart) {
			if !g.isStrike(g.rFrameStart + 1) {
				if g.strikeBonus(g.rFrameStart) > pinsPerFrame {
					return ErrPinCountExceedsPinsOnTheLane
				}
			}
			if b := g.strikeBonus(g.rFrameStart); b > pinsPerFrame && b < 2*pinsPerFrame {
				if !g.isStrike(g.rFrameStart+1) && !g.isStrike(g.rFrameStart+2) {
					return ErrPinCountExceedsPinsOnTheLane
				}
			}
		} else if !g.isSpare(g.rFrameStart) {
			return ErrCannotRollAfterGameOver
		}
		g.completeTheFrame()
	}
	return nil
}

func (g *Game) Score() (int, error) {
	if g.completedFrames() != framesPerGame {
		return 0, ErrPrematureScore
	}
	score := 0
	frameStart := 0
	for frame := 0; frame < framesPerGame; frame++ {
		switch {
		case g.isStrike(frameStart):
			score += pinsPerFrame + g.strikeBonus(frameStart)
			frameStart++
		case g.isSpare(frameStart):
			score += pinsPerFrame + g.spareBonus(frameStart)
			frameStart += maxRollsPerFrame
		default:
			score += g.rawFrameScore(frameStart)
			frameStart += maxRollsPerFrame
		}
	}
	return score, nil
}

func (g *Game) rollsThisFrame() int     { return g.nRolls - g.rFrameStart }
func (g *Game) completeTheFrame()       { g.nFrames++; g.rFrameStart = g.nRolls }
func (g *Game) completedFrames() int    { return g.nFrames }
func (g *Game) isStrike(f int) bool     { return g.rolls[f] == pinsPerFrame }
func (g *Game) rawFrameScore(f int) int { return g.rolls[f] + g.rolls[f+1] }
func (g *Game) spareBonus(f int) int    { return g.rolls[f+2] }
func (g *Game) strikeBonus(f int) int   { return g.rolls[f+1] + g.rolls[f+2] }
func (g *Game) isSpare(f int) bool      { return (g.rolls[f] + g.rolls[f+1]) == pinsPerFrame }
