package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"snakehero/game"
	"testing"
	"time"
)

func TestAll(t *testing.T) {

	//moves := newMoves()
	g := game.GameRequest{}
	g.Board.Width = 11
	g.Board.Height = 11

	snake := [][]int{
		{0, 5}, {1, 5}, {2, 5}, {3, 5}, {4, 5}, {5, 5}, {5, 4}, {5, 3},
	}
	game.AddSTestSnake(&g, snake)

	snake = [][]int{
		{0, 8}, {1, 8}, {2, 8}, {3, 8}, {4, 8}, {5, 8},
	}
	game.AddSTestSnake(&g, snake)

	snake = [][]int{
		//{1, 0}, {0, 0},
		{1, 2}, {1, 1}, {1, 0},
	}
	game.AddTestYou(&g, snake)
	food := []game.Coord{
		//{8, 8}, {7, 7},
		{2, 0}, {5, 5},
	}
	g.Board.Food = food
	//fmt.Println(s)
	board := game.MakeBoard(g)
	x, y, ok := game.FindFood(g.You.Head, board, g.Board.Food)
	assert.Truef(t, ok, "nie ok")
	best := game.FindCoordinates(x, y, g.You.Head)
	fmt.Println("path ", best.Heading)

	moves := game.RankSpace(g.You.Head, board)

	tm := time.Now()
	round := game.Minimax(board, 12, board.GameData.You.ID)
	fmt.Printf("counter %v\n", game.Counter)
	for _, r := range round {
		if r.SnakeId == board.GameData.You.ID {
			fmt.Printf("move: %v\n", r.Move)
		}
	}
	fmt.Printf("%v\n", time.Since(tm))

	fmt.Println(game.FindBest(moves))

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")
}
