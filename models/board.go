package models

import (
	"fmt"
	"github.com/jinzhu/copier"
	"log"
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
	t := make([][]*Tile, game.Board.Height+2)
	for i := range t {
		t[i] = make([]*Tile, game.Board.Width+2)
	}

	board := MyBoard{tiles: t, GameData: &game}

	for _, s := range game.Board.Snakes {
		// todo: Constrictor mode
		for i := 0; i < int(s.Length); i++ {
			t := Tile{X: s.Body[i].X, Y: s.Body[i].Y, board: &board, costIndex: snake, snakeTileNo: i}
			board.SetTile(t.X, t.Y, &t)
		}
	}

	for _, f := range game.Board.Food {
		t := Tile{X: f.X, Y: f.Y, board: &board, costIndex: food}
		board.SetTile(t.X, t.Y, &t)
	}

	for y := 0; y < game.Board.Height; y++ {
		for x := 0; x < game.Board.Width; x++ {
			tile := board.Tile(x, y)
			if tile == nil {
				t := Tile{X: x, Y: y, board: &board, costIndex: empty}
				board.SetTile(t.X, t.Y, &t)
			}
		}
	}

	//walls
	for y := -1; y <= game.Board.Height; y++ {
		t := Tile{X: -1, Y: y, board: &board, costIndex: wall}
		board.SetTile(t.X, t.Y, &t)
		t = Tile{X: game.Board.Width, Y: y, board: &board, costIndex: wall}
		board.SetTile(t.X, t.Y, &t)
	}

	for x := -1; x <= game.Board.Width; x++ {
		t := Tile{X: x, Y: -1, board: &board, costIndex: wall}
		board.SetTile(t.X, t.Y, &t)
		t = Tile{X: x, Y: game.Board.Height, board: &board, costIndex: wall}
		board.SetTile(t.X, t.Y, &t)
	}
	return board
}

func (b *MyBoard) SetTile(x, y int, t *Tile) {
	b.tiles[x+1][y+1] = t
}

func (b *MyBoard) Tile(x, y int) *Tile {
	x++
	y++
	if x < 0 || x > b.GameData.Board.Width+1 || y < 0 || y > b.GameData.Board.Height+1 {
		log.Panicf("##############: tile out of range: %d-1, %d-1\n", x, y)
	}
	return b.tiles[x][y]
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
	tiles := make([][]*Tile, b.GameData.Board.Height+2)
	for i := range tiles {
		tiles[i] = make([]*Tile, b.GameData.Board.Width+2)
	}

	board := MyBoard{tiles: tiles}

	for y, yTiles := range b.tiles {
		for x, t := range yTiles {
			t := Tile{X: t.X, Y: t.Y, board: t.board, costIndex: t.costIndex, snakeTileNo: t.snakeTileNo}
			//board.SetTile(x, y, &t)
			board.tiles[x][y] = &t
		}
	}
	gameRequest := GameRequest{}
	err := copier.Copy(&gameRequest, b.GameData)
	if err != nil {
		panic("cannot copy request")
	}
	board.GameData = &gameRequest
	return board
}

func (b *MyBoard) ApplyMoves(round SnakeMoves) {
	for _, oneMove := range round {
	loop:
		for i := range b.GameData.Board.Snakes {
			s := &b.GameData.Board.Snakes[i]
			if s.ID == oneMove.SnakeId {
				if oneMove.Move.Y == 12 {
					fmt.Println()
				}
				switch b.Tile(oneMove.Move.X, oneMove.Move.Y).costIndex {
				case snake:
					if cId, ok := b.checkH2HCollision(oneMove, s); ok {
						if s.Length <= b.GameData.Board.Snakes[cId].Length {
							s.Health = 0
							if oneMove.SnakeId == b.GameData.You.ID {
								b.GameData.You.Health = 0
							}
							break loop
						}
					}
					b.decriseTileAndHealth(s)
				case wall:
					s.Health = 0
					if oneMove.SnakeId == b.GameData.You.ID {
						b.GameData.You.Health = 0
					}
					break loop
				case food:
					s.Length++
					s.Health = 100
				default:
					b.decriseTileAndHealth(s)
				}

				head := Coord{X: oneMove.Move.X, Y: oneMove.Move.Y}
				body := append([]Coord{}, head)
				body = append(body, s.Body...) //todo: make a copy first?
				s.Body = body
				s.Head = head

				if b.GameData.You.ID == oneMove.SnakeId {
					b.GameData.You = *s
				}

				newHeadTile := Tile{X: oneMove.Move.X, Y: oneMove.Move.Y, board: b}
				b.SetTile(newHeadTile.X, newHeadTile.Y, &newHeadTile)
				break
			}
		}
	}
}

func (b *MyBoard) Clean(round SnakeMoves) {
	for _, oneMove := range round {
		for i := range b.GameData.Board.Snakes {
			if b.GameData.Board.Snakes[i].ID == oneMove.SnakeId {
				if b.GameData.Board.Snakes[i].Health == 0 {
					s := b.GameData.Board.Snakes
					//remove snake
					s[len(s)-1], s[i] = s[i], s[len(s)-1] // swap with last
					b.GameData.Board.Snakes = s[:len(s)-1]
					break
					//if b.GameData.You.ID == oneMove.SnakeId {
					//	b.GameData.You = *s
					//}
				}
			}
		}

	}
}

func (b *MyBoard) decriseTileAndHealth(s *Battlesnake) {
	s.Health--
	lastBodyPart := s.Body[len(s.Body)-1]
	emptyTile := Tile{X: lastBodyPart.X, Y: lastBodyPart.Y, board: b, costIndex: empty}
	b.SetTile(emptyTile.X, emptyTile.Y, &emptyTile)
}

func (b *MyBoard) h2h(oneMove SnakeMove, snake *Battlesnake) bool {
	for i := range b.GameData.Board.Snakes {
		if b.GameData.Board.Snakes[i].Head.X == oneMove.Move.X && b.GameData.Board.Snakes[i].Head.Y == oneMove.Move.Y {
			return snake.Length <= b.GameData.Board.Snakes[i].Length // death
		}
	}
	return true // no H2H
}

func (b *MyBoard) checkH2HCollision(oneMove SnakeMove, snake *Battlesnake) (int, bool) {
	for i := range b.GameData.Board.Snakes {
		if b.GameData.Board.Snakes[i].ID == oneMove.SnakeId {
			continue
		}
		if b.GameData.Board.Snakes[i].Head.X == oneMove.Move.X && b.GameData.Board.Snakes[i].Head.Y == oneMove.Move.Y {
			return i, true
		}
	}
	return 0, false
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
		head := b.Tile(snake.Head.X, snake.Head.Y)
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

func removeMove(moves []Direction, toRemove Direction) []Direction {
	for i, m := range moves {
		if m.Heading == toRemove.Heading {
			moves[len(moves)-1], moves[i] = moves[i], moves[len(moves)-1] // swap with last
			return moves[:len(moves)-1]                                   // truncate last
		}
	}
	return moves
}
