package models

import (
	"github.com/jinzhu/copier"
)

type Direction struct {
	X, Y    int
	Heading string
	Rank    int
}

type MyBoard struct {
	tiles    [][]*Tile
	GameData *GameRequest
}

type SnakeMove struct {
	SnakeId string
	Move    Tile
	Payoff  float64
}

type SnakeMoves []SnakeMove

type NeighbourListWithIterator struct {
	SnakeMoves
	Iterator int
}

type NeighbourTilesForAllSnakes []NeighbourListWithIterator

func MakeBoard(game GameRequest) MyBoard {
	t := make([][]*Tile, game.Board.Height)
	for i := range t {
		t[i] = make([]*Tile, game.Board.Width)
	}

	board := MyBoard{tiles: t, GameData: &game}

	for _, s := range game.Board.Snakes {
		// todo: Constrictor mode
		var i int32
		for i = 0; i < s.Length-1; i++ {
			board.tiles[s.Body[i].X][s.Body[i].Y] =
				&Tile{X: s.Body[i].X, Y: s.Body[i].Y, board: &board, costIndex: snake, snakeTileVanish: int(s.Length - i - 1)}
		}

		//if s.Head.X == play.You.Head.X && s.Head.Y == play.You.Head.Y {
		//	continue
		//}
		//if s.Length > play.You.Length {
		//	for _, m := range NewMoves() {
		//		if m.X < 0 || m.X >= play.Board.Width || m.Y < 0 || m.Y >= play.Board.Height {
		//			continue
		//		}
		//		MyBoard.tiles[s.Head.X+m.X][s.Head.Y+m.Y] = &Tile{X: s.Head.X + m.X, Y: s.Head.Y + m.Y, MyBoard: &MyBoard, costIndex: headAround}
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

func (b *MyBoard) GetTile(x, y int) (*Tile, bool) {
	if x < 0 || x >= b.GameData.Board.Width || y < 0 || y >= b.GameData.Board.Height {
		return nil, false
	}
	return b.tiles[x][y], true
}

func (b *MyBoard) GetBattlesnake(id string) *Battlesnake {
	for i, _ := range b.GameData.Board.Snakes {
		if b.GameData.Board.Snakes[i].ID == id {
			return &b.GameData.Board.Snakes[i]
		}
	}
	panic("battlesnake not found")
}

func (b *MyBoard) CopyBoard() MyBoard {
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
	return MyBoard{tiles: tiles, GameData: &gameRequest}
}

func (b *MyBoard) ApplyMoves(round SnakeMoves) {
	for _, oneMove := range round {
		foundFood := b.tiles[oneMove.Move.X][oneMove.Move.Y].costIndex == food

		newHeadTile := Tile{X: oneMove.Move.X, Y: oneMove.Move.Y, board: b}
		//destiny, ok := b.getTile(newHeadTile.X, newHeadTile.Y)
		//if ok {
		//	if destiny.costIndex
		//} else {
		//
		//}
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
	//// check head2head collision
	//for _, snake1 := range round {
	//	for _, snake2 := range round {
	//		if snake1.SnakeId == snake2.SnakeId {
	//			continue
	//		}
	//		if snake1.Move.X == snake2.Move.X && snake1.Move.Y == snake2.Move.Y {
	//
	//		}
	//	}
	//}
}

func (b *MyBoard) Clean() {

}

func (b *MyBoard) AllCombinations() []SnakeMoves {
	list := b.makeListOfNeighbourTilesForAllSnakes()

	roundList := []SnakeMoves{}
	for {
		round := SnakeMoves{}
		for _, comb := range list {
			round = append(round, comb.SnakeMoves[comb.Iterator])
		}
		if len(round) == 0 {
			return roundList
		}
		roundList = append(roundList, round)
		for i, _ := range list {
			list[i].Iterator++
			if list[i].Iterator < len(list[i].SnakeMoves) {
				break
			}
			list[i].Iterator = 0
		}
		sum := 0
		for _, comb := range list {
			sum += comb.Iterator
		}
		if sum == 0 {
			return roundList
		}
	}

}

func (b *MyBoard) makeListOfNeighbourTilesForAllSnakes() NeighbourTilesForAllSnakes {
	listOfListsOfNeighbours := NeighbourTilesForAllSnakes{}
	for _, snake := range b.GameData.Board.Snakes {
		head, ok := b.GetTile(snake.Head.X, snake.Head.Y)
		if !ok {
			panic("no head in minimax")
		}

		listOfNeighbours := NeighbourListWithIterator{}
		for _, m := range head.Neighbors() {
			move := SnakeMove{SnakeId: snake.ID, Move: *m}
			listOfNeighbours.SnakeMoves = append(listOfNeighbours.SnakeMoves, move)
		}
		if len(listOfNeighbours.SnakeMoves) == 0 {
			return listOfListsOfNeighbours
		}
		listOfListsOfNeighbours = append(listOfListsOfNeighbours, listOfNeighbours)
	}
	return listOfListsOfNeighbours
}
