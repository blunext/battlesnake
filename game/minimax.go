package game

import (
	"github.com/jinzhu/copier"
)

type snakeMove struct {
	SnakeId string
	Move    *Tile
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
	combinations := allCombinations(board)
	for _, round := range combinations {
		newBoard := copyBoard(board) // todo: for next newboard we could revert prev changes
		prepareBoard(board, newBoard, round)
		evaluateRound(newBoard, round) // changes payoff in combination round
		depth--
		if depth == 0 {
			return round
		}
		nextLevel := Minimax(newBoard, depth, heroId)
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
		return round
	}
	panic("ale o co chodzi")
}

func prepareBoard(board board, newBoard board, round snakeMoves) {
	for _, oneMove := range round {
		foundFood := board.tiles[oneMove.Move.x][oneMove.Move.y].costIndex == food

		newHeadPos := Tile{x: oneMove.Move.x, y: oneMove.Move.y, board: oneMove.Move.board}
		newBoard.tiles[oneMove.Move.x][oneMove.Move.y] = &newHeadPos

		for i, snake := range board.GameData.Board.Snakes {
			if snake.ID == oneMove.SnakeId {
				head := Coord{X: oneMove.Move.x, Y: oneMove.Move.y}
				body := append([]Coord{}, head)
				body = append(body, snake.Body...) //todo: make a copy first?
				newBoard.GameData.Board.Snakes[i].Body = body
				newBoard.GameData.Board.Snakes[i].Head = head
				if foundFood {
					newBoard.GameData.Board.Snakes[i].Length++
					newBoard.GameData.Board.Snakes[i].Health = 100
				} else {
					newBoard.GameData.Board.Snakes[i].Health--
					lastBodyPart := snake.Body[len(snake.Body)-1]
					emptyTile := Tile{x: lastBodyPart.X, y: lastBodyPart.Y, board: &newBoard, costIndex: empty}
					newBoard.tiles[lastBodyPart.X][lastBodyPart.Y] = &emptyTile
				}
				if board.GameData.You.ID == oneMove.SnakeId {
					newBoard.GameData.You = newBoard.GameData.Board.Snakes[i]
				}
				break
			}
		}
	}
}

func copyBoard(old board) board {
	tiles := make([][]*Tile, old.GameData.Board.Height)
	for i := range tiles {
		tiles[i] = make([]*Tile, old.GameData.Board.Width)
	}

	for y, yTiles := range old.tiles {
		for x, t := range yTiles {
			tiles[x][y] = &Tile{x: t.x, y: t.y, board: t.board} // todo: snakeTileVanish
		}
	}
	gameRequest := GameRequest{}
	err := copier.Copy(&gameRequest, old.GameData)
	if err != nil {
		panic("cannot copy request")
	}
	return board{tiles: tiles, GameData: &gameRequest}
}

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
			move := snakeMove{SnakeId: snake.ID, Move: m}
			listOfNeighbours.snakeMoves = append(listOfNeighbours.snakeMoves, move)
		}
		listOfListsOfNeighbours = append(listOfListsOfNeighbours, listOfNeighbours)
	}
	return listOfListsOfNeighbours
}
