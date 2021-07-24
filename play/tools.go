package play

import (
	"math/rand"
	"snakehero/models"
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

func AddSTestSnake(g *models.GameRequest, body [][]int) models.Battlesnake {
	snake := models.Battlesnake{}
	for i, pair := range body {
		if i == 0 {
			snake.Head.X = pair[0]
			snake.Head.Y = pair[1]
		}
		snake.Body = append(snake.Body, models.Coord{X: pair[0], Y: pair[1]})
	}
	snake.Length = int32(len(snake.Body))
	snake.ID = randomString(20)
	snake.Health = 100

	g.Board.Snakes = append(g.Board.Snakes, snake)
	return snake
}

func AddTestYou(g *models.GameRequest, body [][]int) {
	snake := AddSTestSnake(g, body)
	g.You = snake
}

//func NewMoves() []models.Direction {
//	return []models.Direction{
//		{0, 1, "up", 0},
//		{0, -1, "down", 0},
//		{-1, 0, "left", 0},
//		{1, 0, "right", 0},
//	}
//}
