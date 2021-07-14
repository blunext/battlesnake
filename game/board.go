package game

type Direction struct {
	X, Y    int
	Heading string
	rank    int
}

type board struct {
	tiles    [][]*Tile
	GameData *GameRequest
}

func MakeBoard(game GameRequest) board {
	t := make([][]*Tile, game.Board.Height)
	for i := range t {
		t[i] = make([]*Tile, game.Board.Width)
	}

	board := board{tiles: t, GameData: &game}

	for _, s := range game.Board.Snakes {
		// todo: Constrictor mode
		var i int32
		for i = 0; i < s.Length-1; i++ {
			board.tiles[s.Body[i].X][s.Body[i].Y] =
				&Tile{x: s.Body[i].X, y: s.Body[i].Y, board: &board, costIndex: snake, snakeTileVanish: int(s.Length - i - 1)}
		}

		//if s.Head.X == game.You.Head.X && s.Head.Y == game.You.Head.Y {
		//	continue
		//}
		//if s.Length > game.You.Length {
		//	for _, m := range NewMoves() {
		//		if m.X < 0 || m.X >= game.Board.Width || m.Y < 0 || m.Y >= game.Board.Height {
		//			continue
		//		}
		//		board.tiles[s.Head.X+m.X][s.Head.Y+m.Y] = &Tile{x: s.Head.X + m.X, y: s.Head.Y + m.Y, board: &board, costIndex: headAround}
		//	}
		//}
	}

	for _, f := range game.Board.Food {
		board.tiles[f.X][f.Y] = &Tile{x: f.X, y: f.Y, board: &board, costIndex: food}
	}

	for y := 0; y < game.Board.Height; y++ {
		for x := 0; x < game.Board.Width; x++ {
			tile := board.tiles[x][y]
			if tile == nil {
				board.tiles[x][y] = &Tile{x: x, y: y, board: &board, costIndex: empty}
			}
		}
	}

	return board
}

func (b *board) getTile(x, y int) (*Tile, bool) {
	if x < 0 || x >= b.GameData.Board.Width || y < 0 || y >= b.GameData.Board.Height {
		return nil, false
	}
	return b.tiles[x][y], true
}

func NewMoves() []Direction {
	return []Direction{
		{0, 1, "up", 0},
		{0, -1, "down", 0},
		{-1, 0, "left", 0},
		{1, 0, "right", 0},
	}
}
