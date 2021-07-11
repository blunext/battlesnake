package game

type headsType struct {
	neighbours []*Tile
	iterator   int
}
type neighboursListType []headsType

type roundType []*Tile
type roundsType []roundType

func Minimax(board board, depth int) {
	//rounds := allCombinations(board)

	//for _, r := range rounds {
	//
	//}

}

func allCombinations(board board) roundsType {
	neighboursList := neighboursListType{}
	for _, snake := range board.gameData.Board.Snakes {
		heads := headsType{}
		head, ok := board.getTile(snake.Head.X, snake.Head.Y)
		if !ok {
			panic("no head in minimax")
		}
		heads.neighbours = head.Neighbors()
		neighboursList = append(neighboursList, heads)
	}

	rounds := roundsType{}
	for {
		round := roundType{}
		for _, comb := range neighboursList {
			round = append(round, comb.neighbours[comb.iterator])
		}
		rounds = append(rounds, round)
		for i, _ := range neighboursList {
			neighboursList[i].iterator++
			if neighboursList[i].iterator < len(neighboursList[i].neighbours) {
				break
			}
			neighboursList[i].iterator = 0
		}
		sum := 0
		for _, comb := range neighboursList {
			sum += comb.iterator
		}
		if sum == 0 {
			break
		}
	}
	return rounds
}
