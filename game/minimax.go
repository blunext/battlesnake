package game

import (
	"math/rand"
	"sync"
)

const MMdepth = 5

type snakeMove struct {
	SnakeId string
	Move    Tile
	payoff  float64
}

type snakeMoves []snakeMove

type neighbourListWithIterator struct {
	snakeMoves
	iterator int
}

type neighbourTilesForAllSnakes []neighbourListWithIterator

type rounds []snakeMoves

const dead = -9999999999999999.0

var Counter int

func Minimax(board board, depth int, heroId string) snakeMoves {

	depth--
	combinations := board.allCombinations()
	if depth == MMdepth-3 {
		movesCh := make(chan snakeMoves)
		wg := sync.WaitGroup{}
		for _, round := range combinations {
			Counter++
			wg.Add(1)
			go func(r snakeMoves, s chan snakeMoves) {
				newBoard := board.copyBoard() // todo: for next newboard we could revert prev changes
				newBoard.applyMoves(round)
				evaluateRound(newBoard, round, heroId) // changes payoff in combination round
				if depth > 0 {
					nextLevel := Minimax(newBoard, depth, heroId)
					mergeLevels(round, nextLevel, heroId)
				}
				s <- round
			}(round, movesCh)
		}
		rounds := rounds{}
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
			newBoard := board.copyBoard() // todo: for next newboard we could revert prev changes
			newBoard.applyMoves(round)
			evaluateRound(newBoard, round, heroId) // changes payoff in combination round
			if depth == 0 {
				return round
			}
			nextLevel := Minimax(newBoard, depth, heroId)
			mergeLevels(round, nextLevel, heroId)
		}
		return bestMove(combinations, heroId)
	}
}

func bestMove(combinations rounds, heroId string) snakeMoves {
	avarage, count := 0.0, 0.0
	for _, round := range combinations {
		for _, r := range round {
			//if r.SnakeId == heroId {
			//	continue
			//}
			avarage += r.payoff
			count++
		}
	}
	avarage = avarage / count
	best := snakeMoves{}
	for _, round := range combinations {
		distance := 0.0
		for _, r := range round {
			if r.SnakeId != heroId {
				if r.payoff-avarage >= distance {
					distance = r.payoff - avarage
					best = round
				}
			}
		}
	}
	if len(best) == 0 {
		if len(combinations) == 0 {
			return snakeMoves{}
		}
		return combinations[rand.Intn(len(combinations))]
	}
	return best
}

func mergeLevels(round snakeMoves, nextLevel snakeMoves, heroId string) {
	for _, r := range round {
		for _, x := range nextLevel {
			if r.SnakeId == heroId && x.SnakeId == heroId && r.payoff < x.payoff {
				r.payoff = x.payoff
				break
			}
			if r.SnakeId == x.SnakeId && r.payoff > x.payoff {
				r.payoff = x.payoff
			}
		}
	}
}
