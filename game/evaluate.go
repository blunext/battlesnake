package game

import "math/rand"

type payoff struct {
	snakeId string
	payoff  float64
}

type payoffs []payoff

// changes payoff in combination round
func evaluateRound(board board, moves snakeMoves) {
	for i, _ := range moves {
		moves[i].payoff = rand.Float64()
	}
}
