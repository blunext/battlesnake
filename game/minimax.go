package game

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

func Minimax(board board, depth int, heroId string) snakeMoves {
	//Counter++
	depth--
	combinations := board.allCombinations()
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

func bestMove(combinations rounds, heroId string) snakeMoves {
	onePlayer := true
	avarage, count := 0.0, 0.0
	for _, round := range combinations {
		for _, r := range round {
			if r.SnakeId == heroId {
				continue
			}
			onePlayer = false
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
		if onePlayer {
			return combinations[0]
		}
		panic("no best moves")
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
