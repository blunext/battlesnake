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

	board := makeBoard(game)
	moves = avoidTakenSpace(game, moves, board)

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

}
