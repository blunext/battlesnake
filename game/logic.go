package game

import (
	"fmt"
	"github.com/beefsack/go-astar"
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

func RankSpace(head Coord, board board) []Direction {
	moves := NewMoves()
	for i, potential := range moves {
		nextMove := head
		nextMove.X += potential.X
		nextMove.Y += potential.Y
		t, present := board.getTile(nextMove.X, nextMove.Y)
		if !present {
			continue
		}
		if t.Cost() < NoPassCost {
			visited := make(map[Coord]*Tile)
			visited[nextMove] = &Tile{}
			moves[i].rank = checkSpace(nextMove, board, 1, visited)
			//fmt.Printf("potential rank %d\n", potential.rank)
		}
	}
	return moves
}

func checkSpace(head Coord, board board, steps int, visited map[Coord]*Tile) int {
	for _, possible := range NewMoves() {
		nextMove := head
		nextMove.X += possible.X
		nextMove.Y += possible.Y
		if _, ok := visited[nextMove]; !ok {
			t, present := board.getTile(nextMove.X, nextMove.Y)
			if !present {
				continue
			}
			if t.Cost() < NoPassCost {
				visited[nextMove] = &Tile{}
				steps++
				steps += checkSpace(nextMove, board, steps, visited)
			}
		}
	}
	return steps
}

func FindBest(moves []Direction) Direction {
	var best Direction
	i := -1
	for _, m := range moves {
		//fmt.Printf("%d, %s\n", m.rank, m.heading)
		if m.rank > i {
			i = m.rank
			best = m
		}
	}
	return best
}

func FindFood(head Coord, board board, food []Coord) (int, int, bool) {
	const maxDistance = 999999999999999999.9
	x, y := -1, -1
	distance := maxDistance
	foodCopied := make([]Coord, len(food))
	copy(foodCopied, food)
	headTile, ok := board.getTile(head.X, head.Y)
	if !ok {
		panic("no head................")
	}
	for _, f := range foodCopied {
		toTile, ok := board.getTile(f.X, f.Y)
		if !ok {
			panic("no food................")
		}
		path, dist, found := astar.Path(headTile, toTile)
		if found {
			if distance > dist {
				distance = dist
				x = path[len(path)-2].(*Tile).x
				y = path[len(path)-2].(*Tile).y
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

func FindCoordinates(x, y int, you Coord) Direction {
	for _, m := range NewMoves() {
		if m.X == x-you.X && m.Y == y-you.Y {
			return m
		}
	}
	panic("no moves...........................................................")
}
