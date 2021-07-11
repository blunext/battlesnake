package game

type headsType struct {
	neighbours []*Tile
	iterator   int
}
type neighboursListType []headsType

type roundType []*Tile
type roundsType []roundType

func Minimax(board board, depth int) {
	for _, rounds := range allCombinations(board) {
		newBoard := copyBoard(board)
		for _, move := range rounds {
			newBoard.tiles[move.x][move.y] = &Tile{x: move.x, y: move.y, board: move.board}

			// dodać przełyżania snake w reqescie

		}
	}

}

func copyBoard(old board) board {
	tiles := make([][]*Tile, old.gameData.Board.Height)
	for i := range tiles {
		tiles[i] = make([]*Tile, old.gameData.Board.Width)
	}

	for y, yTiles := range old.tiles {
		for x, t := range yTiles {
			tiles[x][y] = &Tile{x: t.x, y: t.y, board: t.board} // todo: snakeTileVanish
		}
	}
	return board{tiles: tiles, gameData: old.gameData}
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
			return rounds
		}
	}

}
