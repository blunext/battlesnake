package main

type boardRepresentation map[Coord]struct{}

func makeBoard(game GameRequest) boardRepresentation {
	board := make(map[Coord]struct{})

	for x := -1; x <= game.Board.Width; x++ {
		board[Coord{X: x, Y: -1}] = struct{}{}
		board[Coord{X: x, Y: game.Board.Height}] = struct{}{}
	}
	for y := -1; y <= game.Board.Height; y++ {
		board[Coord{X: -1, Y: y}] = struct{}{}
		board[Coord{X: game.Board.Width, Y: y}] = struct{}{}
	}

	for _, snake := range game.Board.Snakes {
		var i int32
		for i = 0; i < snake.Length-1; i++ { // without last element
			board[snake.Body[i]] = struct{}{}
		}
	}
	return board
}

func avoidTakenSpace(game GameRequest, moves movesSet, board boardRepresentation) movesSet {
	resultMoves := copyMoves(moves)
	for _, possible := range moves {
		nextMove := game.You.Head
		nextMove.X += possible.x
		nextMove.Y += possible.y
		if _, ok := board[nextMove]; ok {
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

func removeMove(moves movesSet, toRemove direction) movesSet {
	for i, m := range moves {
		if m.heading == toRemove.heading {
			moves[len(moves)-1], moves[i] = moves[i], moves[len(moves)-1] // swap with last
			return moves[:len(moves)-1]                                   // truncate last
		}
	}
	return moves
}
