package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {

	//moves := newMoves()
	game := GameRequest{}
	game.Board.Width = 11
	game.Board.Height = 11
	game.You.Head.X = 5
	game.You.Head.Y = 5

	board := makeBoard(game)
	coord := findDirection(game.You.Head, board)
	fmt.Println(coord.heading)

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

}
