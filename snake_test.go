package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"starter-snake-go/game"
	"testing"
)

func TestSomething(t *testing.T) {

	//moves := newMoves()
	g := game.GameRequest{}
	g.Board.Width = 11
	g.Board.Height = 11
	//g.You.Head.X = 5
	//g.You.Head.Y = 5

	snake := []pairs{
		{0, 5}, {1, 5}, {2, 5}, {3, 5}, {4, 5}, {5, 5}, {5, 4}, {5, 3},
	}
	addSnake(&g, snake)
	snake = []pairs{
		//{1, 0}, {0, 0},
		{1, 2}, {1, 1}, {1, 1},
	}
	addYou(&g, snake)
	food := []game.Coord{
		//{8, 8}, {7, 7},
		{2, 0}, {5, 5},
	}
	g.Board.Food = food
	//fmt.Println(s)
	board := game.MakeBoard(g)
	c, ok := game.FindFood(g.You.Head, board, g.Board.Food)
	assert.Truef(t, ok, "nie ok")
	best := game.FindCoordinates(c, g.You.Head)
	fmt.Println("path ", best.Heading)

	moves := game.RankSpace(g.You.Head, board)

	fmt.Println(game.FindBest(moves))

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

}
