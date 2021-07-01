package main

import (
	"fmt"
	"github.com/BattlesnakeOfficial/starter-snake-go/mygame"
	"github.com/stretchr/testify/assert"
	"testing"
)

type pairs []int

func addSnake(game *mygame.GameRequest, body []pairs) mygame.Battlesnake {
	snake := mygame.Battlesnake{}
	for i, pair := range body {
		if i == 0 {
			snake.Head.X = pair[0]
			snake.Head.Y = pair[1]
		}
		snake.Body = append(snake.Body, mygame.Coord{X: pair[0], Y: pair[1]})
	}
	game.Board.Snakes = append(game.Board.Snakes, snake)
	return snake
}

func addYou(game *mygame.GameRequest, body []pairs) {
	snake := addSnake(game, body)
	game.You = snake
}
func TestSomething(t *testing.T) {

	//moves := newMoves()
	game := mygame.GameRequest{}
	game.Board.Width = 11
	game.Board.Height = 11
	game.You.Head.X = 5
	game.You.Head.Y = 5

	snake := []pairs{
		{0, 3}, {1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3}, {5, 2}, {5, 1},
	}
	addSnake(&game, snake)
	snake = []pairs{
		{1, 0}, {0, 0},
	}
	addYou(&game, snake)
	//fmt.Println(s)
	board := mygame.makeBoard(game)
	moves := mygame.rankSpace(game.You.Head, board)

	fmt.Println(mygame.findBest(moves))

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

}
