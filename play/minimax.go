package play

import (
	"math/rand"
	"snakehero/models"
	"sync"
)

const MMdepth = 4

var Counter int

func Minimax(board models.MyBoard, depth int, heroId string) models.SnakeMoves {

	depth--
	combinations := board.AllCombinations()
	if depth == MMdepth-33 {
		movesCh := make(chan models.SnakeMoves)
		wg := sync.WaitGroup{}
		for _, round := range combinations {
			Counter++
			wg.Add(1)
			go func(r models.SnakeMoves, moves chan models.SnakeMoves) {
				newBoard := board.CopyBoard() // todo: for next newboard we could revert prev changes
				newBoard.ApplyMoves(r)
				evaluateRound(newBoard, r, heroId) // changes payoff in combination round
				newBoard.Clean(r)                  // removes dead snakes
				//jeśeli jakiś snake pozostał chyba że jesteśmy sami?
				if depth > 0 {
					nextLevel := Minimax(newBoard, depth, heroId)
					mergeLevels(r, nextLevel, heroId)
				}
				moves <- r
			}(round, movesCh)
		}
		rounds := []models.SnakeMoves{}
		go func() {
			for r := range movesCh {
				rounds = append(rounds, r)
				wg.Done()
			}
		}()
		wg.Wait()
		return bestMove(rounds, heroId)
	} else {
		for _, round := range combinations {
			newBoard := board.CopyBoard() // todo: for next newboard we could revert prev changes
			newBoard.ApplyMoves(round)
			evaluateRound(newBoard, round, heroId) // changes payoff in combination round
			newBoard.Clean(round)
			if depth == 0 {
				return round
			}

			nextLevel := Minimax(newBoard, depth, heroId)
			mergeLevels(round, nextLevel, heroId)
		}
		return bestMove(combinations, heroId)
	}
}

func bestMove(combinations []models.SnakeMoves, heroId string) models.SnakeMoves {
	avarage, count := 0.0, 0.0
	for _, round := range combinations {
		for _, r := range round {
			//if r.SnakeId == heroId {
			//	continue
			//}
			avarage += r.Payoff
			count++
		}
	}
	avarage = avarage / count
	best := models.SnakeMoves{}
	for _, round := range combinations {
		distance := 0.0
		for _, r := range round {
			if r.SnakeId != heroId {
				if r.Payoff-avarage >= distance {
					distance = r.Payoff - avarage
					best = round
				}
			}
		}
	}
	if len(best) == 0 {
		if len(combinations) == 0 {
			return models.SnakeMoves{}
		}
		return combinations[rand.Intn(len(combinations))]
	}
	return best
}

func mergeLevels(round models.SnakeMoves, nextLevel models.SnakeMoves, heroId string) {
	for _, r := range round {
		for _, x := range nextLevel {
			if r.SnakeId == heroId && x.SnakeId == heroId && r.Payoff < x.Payoff {
				r.Payoff = x.Payoff
				break
			}
			if r.SnakeId == x.SnakeId && r.Payoff > x.Payoff {
				r.Payoff = x.Payoff
			}
		}
	}
}
