package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {

	moves := newMoves()
	game := GameRequest{}
	game.Board.Width = 11
	game.Board.Height = 11
	game.You.Head.X = 0
	game.You.Head.X = 0

	moves = avoidBoundaries(game, moves)

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

}
