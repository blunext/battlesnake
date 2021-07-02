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
//		nextMove.X += possible.x
//		nextMove.Y += possible.y
//		if _, ok := board[nextMove]; ok {
//			resultMoves = removeMove(resultMoves, possible)
//		}
//	}
//	return resultMoves
//}

//func removeMove(moves movesSet, toRemove direction) movesSet {
//	for i, m := range moves {
//		if m.Heading == toRemove.Heading {
//			moves[len(moves)-1], moves[i] = moves[i], moves[len(moves)-1] // swap with last
//			return moves[:len(moves)-1]                                   // truncate last
//		}
//	}
//	return moves
//}

func RankSpace(head Coord, board coordinatesMap) []direction {
	moves := newMoves()
	for i, potential := range moves {
		nextMove := head
		nextMove.X += potential.x
		nextMove.Y += potential.y
		t, present := board[nextMove]
		if !present || t.Cost() < NoPassCost {
			visited := make(map[Coord]Tile)
			visited[nextMove] = Tile{}
			moves[i].rank = checkSpace(nextMove, board, 1, visited)
			//fmt.Printf("potential rank %d\n", potential.rank)
		}
	}
	return moves
}

func checkSpace(head Coord, board coordinatesMap, steps int, visited coordinatesMap) int {
	for _, possible := range newMoves() {
		nextMove := head
		nextMove.X += possible.x
		nextMove.Y += possible.y
		if _, ok := visited[nextMove]; !ok {
			t, present := board[nextMove]
			if !present || t.Cost() < NoPassCost {
				visited[nextMove] = Tile{}
				steps++
				steps += checkSpace(nextMove, board, steps, visited)
			}
		}
	}
	return steps
}

func FindBest(moves []direction) direction {
	var best direction
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

func FindFood(head Coord, board coordinatesMap, food []Coord) {
	foodCopied := make([]Coord, len(food))
	copy(foodCopied, food)
	headTile := board[head]
	for _, f := range foodCopied {
		toTile := board[f]
		_, distance, found := astar.Path(&headTile, &toTile)
		if found {
			fmt.Printf("distance %v\n", distance)
		} else {
			fmt.Printf("distance not found\n")

		}
	}
}
