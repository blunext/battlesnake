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
		for i = 0; i < snake.Length-1; i++ {
			board[snake.Body[i]] = struct{}{}
		}
		if snake.Head.X == game.You.Head.X && snake.Head.Y == game.You.Head.Y {
			continue
		}

		if snake.Length > game.You.Length {
			for _, m := range newMoves() {
				board[Coord{X: snake.Head.X + m.x, Y: snake.Head.Y + m.y}] = struct{}{}
			}
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

func rankSpace(head Coord, board coordinatesMap) []direction {
	moves := newMoves()
	for i, potential := range moves {
		nextMove := head
		nextMove.X += potential.x
		nextMove.Y += potential.y
		if _, ok := board[nextMove]; !ok {
			visited := make(map[Coord]struct{})
			visited[nextMove] = struct{}{}
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
			if _, ok = board[nextMove]; !ok {
				visited[nextMove] = struct{}{}
				steps++
				steps += checkSpace(nextMove, board, steps, visited)
			}
		}
	}
	return steps
}

func findBest(moves []direction) direction {
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
