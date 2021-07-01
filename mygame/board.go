package mygame

type Game struct {
	ID      string `json:"id"`
	Timeout int32  `json:"timeout"`
}

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Battlesnake struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Health  int32   `json:"health"`
	Body    []Coord `json:"body"`
	Latency string  `json:latency`
	Head    Coord   `json:"head"`
	Length  int32   `json:"length"`
	Shout   string  `json:"shout"`
}

type Board struct {
	Height int           `json:"height"`
	Width  int           `json:"width"`
	Food   []Coord       `json:"food"`
	Snakes []Battlesnake `json:"snakes"`
}

type BattlesnakeInfoResponse struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}

type GameRequest struct {
	Game  Game        `json:"game"`
	Turn  int         `json:"turn"`
	Board Board       `json:"board"`
	You   Battlesnake `json:"you"`
}

type MoveResponse struct {
	Move  string `json:"move"`
	Shout string `json:"shout,omitempty"`
}

type direction struct {
	x, y    int
	Heading string
	rank    int
}

type movesSet []direction

type tile struct {
	*Coord
	board coordinatesMap
}

type coordinatesMap map[Coord]tile

//func (t *tile) PathNeighbors() []astar.Pather {
//	var neighbors []astar.Pather
//	for _, next := range newMoves() {
//		c := Coord{X: t.X + next.x, Y: t.Y + next.y}
//		if tile, ok := t.board[c]; ok {
//			neighbors = append(neighbors, &tile)
//		}
//	}
//	return neighbors
//}
//
//func (t *tile) PathNeighborCost(to astar.Pather) float64 {
//	return t.MovementCost
//}
//
//func (t *tile) PathEstimatedCost(to astar.Pather) float64 {
//	return t.ManhattanDistance(to)
//}

func MakeBoard(game GameRequest) coordinatesMap {
	board := make(map[Coord]tile)

	for x := -1; x <= game.Board.Width; x++ {
		c := Coord{X: x, Y: -1}
		board[c] = tile{Coord: &c, board: board}
		c = Coord{X: x, Y: game.Board.Height}
		board[c] = tile{Coord: &c, board: board}
	}
	for y := -1; y <= game.Board.Height; y++ {
		c := Coord{X: -1, Y: y}
		board[c] = tile{Coord: &c, board: board}
		c = Coord{X: game.Board.Width, Y: y}
		board[c] = tile{Coord: &c, board: board}
	}

	for _, snake := range game.Board.Snakes {
		// TODO: tail dispears when no food is consumed
		var i int32
		for i = 0; i < snake.Length-1; i++ {
			board[snake.Body[i]] = tile{Coord: &snake.Body[i], board: board}
		}
		if snake.Head.X == game.You.Head.X && snake.Head.Y == game.You.Head.Y {
			continue
		}

		if snake.Length > game.You.Length {
			for _, m := range newMoves() {
				c := Coord{X: snake.Head.X + m.x, Y: snake.Head.Y + m.y}
				board[c] = tile{Coord: &c, board: board}
			}
		}
	}
	return board
}

func newMoves() movesSet {
	return movesSet{
		{0, 1, "up", 0},
		{0, -1, "down", 0},
		{-1, 0, "left", 0},
		{1, 0, "right", 0},
	}
}
