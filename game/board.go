package game

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
	Latency string  `json:"latency"`
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
type coordinatesMap map[Coord]Tile

func MakeBoard(game GameRequest) coordinatesMap {
	board := make(map[Coord]Tile)

	for x := -1; x <= game.Board.Width; x++ {
		c := Coord{X: x, Y: -1}
		board[c] = Tile{Coord: &c, board: board, costIndex: wall}
		c = Coord{X: x, Y: game.Board.Height}
		board[c] = Tile{Coord: &c, board: board, costIndex: wall}
	}
	for y := -1; y <= game.Board.Height; y++ {
		c := Coord{X: -1, Y: y}
		board[c] = Tile{Coord: &c, board: board, costIndex: wall}
		c = Coord{X: game.Board.Width, Y: y}
		board[c] = Tile{Coord: &c, board: board, costIndex: wall}
	}

	for _, s := range game.Board.Snakes {
		// TODO: tail dispears when no food is consumed
		var i int32
		for i = 0; i < s.Length-1; i++ {
			board[s.Body[i]] = Tile{Coord: &s.Body[i], board: board, costIndex: snake}
		}
		if s.Head.X == game.You.Head.X && s.Head.Y == game.You.Head.Y {
			continue
		}

		if s.Length > game.You.Length {
			for _, m := range newMoves() {
				c := Coord{X: s.Head.X + m.x, Y: s.Head.Y + m.y}
				board[c] = Tile{Coord: &c, board: board, costIndex: headAround}
			}
		}
	}
	for _, f := range game.Board.Food {
		board[f] = Tile{Coord: &f, board: board, costIndex: food}
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
