package main

import "snakehero/game"

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
	snake.Length = int32(len(snake.Body))
	g.Board.Snakes = append(g.Board.Snakes, snake)
	return snake
}

func addYou(g *game.GameRequest, body []pairs) {
	snake := addSnake(g, body)
	g.You = snake
}
