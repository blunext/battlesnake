package game

import (
	"fmt"
	"github.com/beefsack/go-astar"
)

//func avoidTakenSpace(head Coord, moves movesSet, board coordinatesMap) movesSet {
//	resultMoves := make(movesSet, len(moves))
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

//func removeMove(moves movesSet, toRemove Direction) movesSet {
//	for i, m := range moves {
//		if m.Heading == toRemove.Heading {
//			moves[len(moves)-1], moves[i] = moves[i], moves[len(moves)-1] // swap with last
//			return moves[:len(moves)-1]                                   // truncate last
//		}
//	}
//	return moves
//}

func RankSpace(head Coord, board coordinatesMap) []Direction {
	moves := NewMoves()
	for i, potential := range moves {
		nextMove := head
		nextMove.X += potential.X
		nextMove.Y += potential.Y
		t, present := board[nextMove]
		if !present || t.Cost() < NoPassCost {
			visited := make(map[Coord]*Tile)
			visited[nextMove] = &Tile{}
			moves[i].rank = checkSpace(nextMove, board, 1, visited)
			//fmt.Printf("potential rank %d\n", potential.rank)
		}
	}
	return moves
}

func checkSpace(head Coord, board coordinatesMap, steps int, visited coordinatesMap) int {
	for _, possible := range NewMoves() {
		nextMove := head
		nextMove.X += possible.X
		nextMove.Y += possible.Y
		if _, ok := visited[nextMove]; !ok {
			t, present := board[nextMove]
			if !present || t.Cost() < NoPassCost {
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

func FindFood(head Coord, board coordinatesMap, food []Coord) (Coord, bool) {
	const maxDistance = 999999999999999999.9
	distance := maxDistance
	foodCopied := make([]Coord, len(food))
	copy(foodCopied, food)
	headTile := board[head]
	nextMove := &Coord{}
	for _, f := range foodCopied {
		toTile := board[f]
		path, dist, found := astar.Path(headTile, toTile)
		if found {
			if distance > dist {
				distance = dist
				nextMove = path[len(path)-2].(*Tile).Coord
			}
			//fmt.Printf("path %v -> %v\n", head, f)
			//fmt.Printf("dist %v\n", dist)
			//fmt.Printf("path: ")
			//for _, v := range path {
			//	fmt.Printf("%v, ", v.(*Tile).Coord)
			//}
			//fmt.Println()
		} else {
			fmt.Printf("cannot find the path\n")
		}
	}
	if distance == maxDistance {
		return *nextMove, false
	}
	return *nextMove, true
}

func FindCoordinates(c Coord, you Coord) Direction {
	for _, m := range NewMoves() {
		if m.X == c.X-you.X && m.Y == c.Y-you.Y {
			return m
		}
	}
	panic("no moves")
}
