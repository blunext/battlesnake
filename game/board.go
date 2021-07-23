package game

import (
	"github.com/jinzhu/copier"
)

type Direction struct {
	X, Y    int
	Heading string
	rank    int
}

type board struct {
	tiles    [][]*Tile
	GameData *GameRequest
}

func MakeBoard(game GameRequest) board {
	t := make([][]*Tile, game.Board.Height)
	for i := range t {
		t[i] = make([]*Tile, game.Board.Width)
	}

	board := board{tiles: t, GameData: &game}

	for _, s := range game.Board.Snakes {
		// todo: Constrictor mode
		var i int32
		for i = 0; i < s.Length-1; i++ {
			board.tiles[s.Body[i].X][s.Body[i].Y] =
				&Tile{X: s.Body[i].X, Y: s.Body[i].Y, board: &board, costIndex: snake, snakeTileVanish: int(s.Length - i - 1)}
		}

		//if s.Head.X == game.You.Head.X && s.Head.Y == game.You.Head.Y {
		//	continue
		//}
		//if s.Length > game.You.Length {
		//	for _, m := range NewMoves() {
		//		if m.X < 0 || m.X >= game.Board.Width || m.Y < 0 || m.Y >= game.Board.Height {
		//			continue
		//		}
		//		board.tiles[s.Head.X+m.X][s.Head.Y+m.Y] = &Tile{X: s.Head.X + m.X, Y: s.Head.Y + m.Y, board: &board, costIndex: headAround}
		//	}
		//}
	}

	for _, f := range game.Board.Food {
		board.tiles[f.X][f.Y] = &Tile{X: f.X, Y: f.Y, board: &board, costIndex: food}
	}

	for y := 0; y < game.Board.Height; y++ {
		for x := 0; x < game.Board.Width; x++ {
			tile := board.tiles[x][y]
			if tile == nil {
				board.tiles[x][y] = &Tile{X: x, Y: y, board: &board, costIndex: empty}
			}
		}
	}

	return board
}

func (b *board) getTile(x, y int) (*Tile, bool) {
	if x < 0 || x >= b.GameData.Board.Width || y < 0 || y >= b.GameData.Board.Height {
		return nil, false
	}
	return b.tiles[x][y], true
}

func (b *board) getBattlesnake(id string) *Battlesnake {
	for i, _ := range b.GameData.Board.Snakes {
		if b.GameData.Board.Snakes[i].ID == id {
			return &b.GameData.Board.Snakes[i]
		}
	}
	panic("battlesnake not found")
}

func (b *board) copyBoard() board {
	tiles := make([][]*Tile, b.GameData.Board.Height)
	for i := range tiles {
		tiles[i] = make([]*Tile, b.GameData.Board.Width)
	}

	for y, yTiles := range b.tiles {
		for x, t := range yTiles {
			tiles[x][y] = &Tile{X: t.X, Y: t.Y, board: t.board, costIndex: t.costIndex} // todo: snakeTileVanish
		}
	}
	gameRequest := GameRequest{}
	err := copier.Copy(&gameRequest, b.GameData)
	if err != nil {
		panic("cannot copy request")
	}
	return board{tiles: tiles, GameData: &gameRequest}
}

func (b *board) applyMoves(round snakeMoves) {
	for _, oneMove := range round {
		foundFood := b.tiles[oneMove.Move.X][oneMove.Move.Y].costIndex == food

		newHeadTile := Tile{X: oneMove.Move.X, Y: oneMove.Move.Y, board: b}
		b.tiles[oneMove.Move.X][oneMove.Move.Y] = &newHeadTile

		for i := range b.GameData.Board.Snakes {
			snake := &b.GameData.Board.Snakes[i]
			if snake.ID == oneMove.SnakeId {
				head := Coord{X: oneMove.Move.X, Y: oneMove.Move.Y}
				body := append([]Coord{}, head)
				body = append(body, snake.Body...) //todo: make a copy first?
				snake.Body = body
				snake.Head = head
				if foundFood {
					snake.Length++
					snake.Health = 100
				} else {
					snake.Health--
					lastBodyPart := snake.Body[len(snake.Body)-1]
					emptyTile := Tile{X: lastBodyPart.X, Y: lastBodyPart.Y, board: b, costIndex: empty}
					b.tiles[lastBodyPart.X][lastBodyPart.Y] = &emptyTile
				}
				if b.GameData.You.ID == oneMove.SnakeId {
					b.GameData.You = *snake
				}
				break
			}
		}
	}
}

func (b *board) allCombinations() rounds {
	list := b.makeListOfNeighbourTilesForAllSnakes()

	roundList := rounds{}
	for {
		round := snakeMoves{}
		for _, comb := range list {
			round = append(round, comb.snakeMoves[comb.iterator])
		}
		if len(round) == 0 {
			return roundList
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

func (b *board) makeListOfNeighbourTilesForAllSnakes() neighbourTilesForAllSnakes {
	listOfListsOfNeighbours := neighbourTilesForAllSnakes{}
	for _, snake := range b.GameData.Board.Snakes {
		head, ok := b.getTile(snake.Head.X, snake.Head.Y)
		if !ok {
			panic("no head in minimax")
		}

		listOfNeighbours := neighbourListWithIterator{}
		for _, m := range head.Neighbors() {
			move := snakeMove{SnakeId: snake.ID, Move: *m}
			listOfNeighbours.snakeMoves = append(listOfNeighbours.snakeMoves, move)
		}
		if len(listOfNeighbours.snakeMoves) == 0 {
			return listOfListsOfNeighbours
		}
		listOfListsOfNeighbours = append(listOfListsOfNeighbours, listOfNeighbours)
	}
	return listOfListsOfNeighbours
}

//func (b *board) evaluateRound(moves []snakeMove, heroId string) {
//	for i := range moves {
//		moves[i].payoff = float64(board.getBattlesnake(moves[i].SnakeId).Length)
//		if moves[i].payoff == 4 || moves[i].payoff == 9 {
//			fmt.Println("sdsdddaaaaaaaaaaa")
//		}
//	}
//
//}
