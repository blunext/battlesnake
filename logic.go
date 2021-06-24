package main

type coordinatesMap map[Coord]struct{}

func makeBoard(game GameRequest) coordinatesMap {
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
		// TODO: tail dispears when no food is consumed
		var i int32
		for i = 0; i < snake.Length; i++ {
			board[snake.Body[i]] = struct{}{}
		}

	}
	return board
}

func avoidTakenSpace(head Coord, moves movesSet, board coordinatesMap) movesSet {
	resultMoves := copyMoves(moves)
	for _, possible := range moves {
		nextMove := head
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

func findDirection(head Coord, board coordinatesMap) direction {
	num := -1
	moves := newMoves()
	winner := moves[0]
	for _, possible := range moves {
		nextMove := head
		nextMove.X += possible.x
		nextMove.Y += possible.y
		if _, ok := board[nextMove]; !ok {
			checked := make(map[Coord]struct{})
			checked[nextMove] = struct{}{}
			count := checkSpace(nextMove, board, 1, checked)
			if num < count {
				winner = possible
				num = count
			}
		}
	}

	return winner
}

func checkSpace(head Coord, board coordinatesMap, steps int, checked coordinatesMap) int {
	for _, possible := range newMoves() {
		nextMove := head
		nextMove.X += possible.x
		nextMove.Y += possible.y
		if _, ok := board[nextMove]; !ok {
			if _, ok = checked[nextMove]; !ok {
				checked[nextMove] = struct{}{}
				steps++
				steps += checkSpace(nextMove, board, steps, checked)
			}
		}
	}
	return steps
}
