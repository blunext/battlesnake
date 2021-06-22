package main

func avoidBoundaries(game GameRequest, moves movesSet) movesSet {
	resultMoves := copyMoves(moves)
	for _, possible := range moves {
		nextMove := game.You.Head
		nextMove.X += possible.x
		nextMove.Y += possible.y
		if nextMove.X < 0 || nextMove.X >= game.Board.Width || nextMove.Y < 0 || nextMove.Y >= game.Board.Height {
			resultMoves = removeMove(resultMoves, possible)
		}
	}
	return resultMoves
}

func copyMoves(moves movesSet) movesSet {
	var resultMoves movesSet
	for _, move := range moves {
		resultMoves = append(resultMoves, move)
	}
	return resultMoves
}

func avoidSelf(game GameRequest, moves movesSet) movesSet {
	resultMoves := copyMoves(moves)
	for _, possible := range moves {
		nextMove := game.You.Head
		nextMove.X += possible.x
		nextMove.Y += possible.y
		for _, coord := range game.You.Body {
			if nextMove.X == coord.X && nextMove.Y == coord.Y {
				resultMoves = removeMove(resultMoves, possible)
			}
		}
	}
	return resultMoves
}

func removeMove(moves movesSet, toRemove direction) movesSet {
	for i, m := range moves {
		if m.heading == toRemove.heading {
			moves[len(moves)-1], moves[i] = moves[i], moves[len(moves)-1] // swap with last
			return moves[:len(moves)-1]                                   // truncate last
		}
	}
	return moves
}
