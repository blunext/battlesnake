package game

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func AddSTestSnake(g *GameRequest, body [][]int) Battlesnake {
	snake := Battlesnake{}
	for i, pair := range body {
		if i == 0 {
			snake.Head.X = pair[0]
			snake.Head.Y = pair[1]
		}
		snake.Body = append(snake.Body, Coord{X: pair[0], Y: pair[1]})
	}
	snake.Length = int32(len(snake.Body))
	snake.ID = randomString(20)
	snake.Health = 100

	g.Board.Snakes = append(g.Board.Snakes, snake)
	return snake
}

func AddTestYou(g *GameRequest, body [][]int) {
	snake := AddSTestSnake(g, body)
	g.You = snake
}
