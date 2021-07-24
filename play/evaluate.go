package play

import "snakehero/models"

const dead = -9999999.0

// changes payoff in combination round
func evaluateRound(board models.MyBoard, moves []models.SnakeMove, heroId string) {
	sign := 1
	for i := range moves {
		snake := board.GetBattlesnake(moves[i].SnakeId)
		if snake.ID == heroId {
			sign = 1
		} else {
			sign = -1
		}
		switch snake.Health {
		case 0:
			moves[i].Payoff = dead
		default:
			moves[i].Payoff = float64(int(snake.Health) / 100 * sign)
		}
	}
}
