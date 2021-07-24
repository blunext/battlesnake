package play

import (
	"fmt"
	"github.com/beefsack/go-astar"
	"snakehero/models"
)

//func avoidTakenSpace(head Coord, moves []Direction, board boardDefinition) []Direction {
//	resultMoves := make([]Direction, len(moves))
//	copy(resultMoves, moves)
//	for _, possible := range moves {
//		nextMove := head
//		nextMove.X += possible.X
//		nextMove.Y += possible.Y
//		if _, ok := board[nextMove]; ok {
//			resultMoves = removeMove(resultMoves, possible)
//		}
//	}
//	return resultMoves
//}
//
//func removeMove(moves []Direction, toRemove Direction) []Direction {
//	for i, m := range moves {
//		if m.Heading == toRemove.Heading {
//			moves[len(moves)-1], moves[i] = moves[i], moves[len(moves)-1] // swap with last
//			return moves[:len(moves)-1]                                   // truncate last
//		}
//	}
//	return moves
//}

func RankSpace(head models.Coord, board models.MyBoard) []models.Direction {
	moves := models.NewMoves()
	for i, potential := range moves {
		nextMove := head
		nextMove.X += potential.X
		nextMove.Y += potential.Y
		t := board.Tile(nextMove.X, nextMove.Y)

		if t.Cost() < models.NoPassCost {
			visited := make(map[models.Coord]*models.Tile)
			visited[nextMove] = &models.Tile{}
			moves[i].Rank = checkSpace(nextMove, board, 1, visited)
			//fmt.Printf("potential rank %d\n", potential.rank)
		}
	}
	return moves
}

func checkSpace(head models.Coord, board models.MyBoard, steps int, visited map[models.Coord]*models.Tile) int {
	for _, possible := range models.NewMoves() {
		nextMove := head
		nextMove.X += possible.X
		nextMove.Y += possible.Y
		if _, ok := visited[nextMove]; !ok {
			t := board.Tile(nextMove.X, nextMove.Y)
			if t.Cost() < models.NoPassCost {
				visited[nextMove] = &models.Tile{}
				steps++
				steps += checkSpace(nextMove, board, steps, visited)
			}
		}
	}
	return steps
}

func FindBest(moves []models.Direction) models.Direction {
	var best models.Direction
	i := -1
	for _, m := range moves {
		//fmt.Printf("%d, %s\n", m.rank, m.heading)
		if m.Rank > i {
			i = m.Rank
			best = m
		}
	}
	return best
}

func FindFood(head models.Coord, board models.MyBoard, food []models.Coord) (int, int, bool) {
	const maxDistance = 999999999999999999.9
	x, y := -1, -1
	distance := maxDistance
	foodCopied := make([]models.Coord, len(food))
	copy(foodCopied, food)
	headTile := board.Tile(head.X, head.Y)
	for _, f := range foodCopied {
		toTile := board.Tile(f.X, f.Y)
		path, dist, found := astar.Path(headTile, toTile)
		if found {
			if distance > dist {
				distance = dist
				x = path[len(path)-2].(*models.Tile).X
				y = path[len(path)-2].(*models.Tile).Y
			}
		} else {
			fmt.Printf("cannot find the path\n")
		}
	}
	if distance == maxDistance {
		return x, y, false
	}
	return x, y, true
}

func FindCoordinates(x, y int, you models.Coord) models.Direction {
	for _, m := range models.NewMoves() {
		if m.X == x-you.X && m.Y == y-you.Y {
			return m
		}
	}
	panic("no moves...........................................................")
}
