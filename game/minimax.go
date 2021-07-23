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
		distanse := 0.0
		for _, r := range round {
			if r.SnakeId != heroId {
				if r.payoff-avarage >= distanse {
					distanse = r.payoff - avarage
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
			if r.SnakeId == heroId && r.SnakeId == x.SnakeId && r.payoff < x.payoff {
				r.payoff = x.payoff
				break
			}
			if r.SnakeId == x.SnakeId && r.payoff > x.payoff {
				r.payoff = x.payoff
			}
		}
	}
}

/*
func allCombinations(board board) rounds {
	list := makeListOfNeighbourTilesForAllSnakes(board)

	roundList := rounds{}
	for {
		round := snakeMoves{}
		for _, comb := range list {
			round = append(round, comb.snakeMoves[comb.iterator])
		}
		roundList = append(roundList, round)
		for i, _ := range list {
			list[i].iterator++
			if list[i].iterator < len(list[i].snakeMoves) {
				break
			}
			list[i].iterator = 0
		}
		sum := 0
		for _, comb := range list {
			sum += comb.iterator
		}
		if sum == 0 {
			return roundList
		}
	}

}

func makeListOfNeighbourTilesForAllSnakes(board board) neighbourTilesForAllSnakes {
	listOfListsOfNeighbours := neighbourTilesForAllSnakes{}
	for _, snake := range board.GameData.Board.Snakes {
		listOfNeighbours := neighbourListWithIterator{}
		head, ok := board.getTile(snake.Head.X, snake.Head.Y)
		if !ok {
			panic("no head in minimax")
		}

		for _, m := range head.Neighbors() {
			move := snakeMove{SnakeId: snake.ID, Move: *m}
			listOfNeighbours.snakeMoves = append(listOfNeighbours.snakeMoves, move)
		}
		listOfListsOfNeighbours = append(listOfListsOfNeighbours, listOfNeighbours)
	}
	return listOfListsOfNeighbours
}
*/
