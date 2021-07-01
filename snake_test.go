package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"starter-snake-go/game"
	"testing"
)

type pairs []int

func addSnake(g *game.GameRequest, body []pairs) game.Battlesnake {
	snake := game.Battlesnake{}
	for i, pair := range body {
		if i == 0 {
			snake.Head.X = pair[0]
			snake.Head.Y = pair[1]
		}
		snake.Body = append(snake.Body, game.Coord{X: pair[0], Y: pair[1]})
	}
	g.Board.Snakes = append(g.Board.Snakes, snake)
	return snake
}

func addYou(g *game.GameRequest, body []pairs) {
	snake := addSnake(g, body)
	g.You = snake
}
func TestSomething(t *testing.T) {

	//moves := newMoves()
	g := game.GameRequest{}
	g.Board.Width = 11
	g.Board.Height = 11
	g.You.Head.X = 5
	g.You.Head.Y = 5

	snake := []pairs{
		{0, 3}, {1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3}, {5, 2}, {5, 1},
	}
	addSnake(&g, snake)
	snake = []pairs{
		{1, 0}, {0, 0},
	}
	addYou(&g, snake)
	//fmt.Println(s)
	board := game.MakeBoard(g)
	moves := game.RankSpace(g.You.Head, board)

	fmt.Println(game.FindBest(moves))

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

}
