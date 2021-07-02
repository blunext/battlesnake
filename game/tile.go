package game

import (
	"fmt"
	"github.com/beefsack/go-astar"
)

const (
	empty = iota
	food
	headAround
	snake
	wall
)

const NoPassCost = 99999999

var tileKindCost = map[int]float64{
	empty:      1.0,
	food:       1.0,
	headAround: NoPassCost, // todo: change
	snake:      NoPassCost,
	wall:       NoPassCost,
}

type Tile struct {
	*Coord
	board     coordinatesMap
	costIndex int
}

func (t *Tile) PathNeighbors() []astar.Pather {
	var neighbors []astar.Pather
	for _, next := range newMoves() {
		c := Coord{X: t.X + next.x, Y: t.Y + next.y}
		neighborTile, present := t.board[c]
		switch present {
		case true:
			//fmt.Printf("%d, %d\n", c.X, c.Y)
			neighbors = append(neighbors, &neighborTile)
		case false:
			fmt.Printf("%d, %d\n", c.X, c.Y)
			neighborTile = Tile{Coord: &c, board: t.board, costIndex: empty}
			t.board[c] = neighborTile
			neighbors = append(neighbors, &neighborTile)
		}
	}
	return neighbors
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	tile := to.(*Tile)
	return tile.Cost()
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

func (t *Tile) Cost() float64 {
	return tileKindCost[t.costIndex]
}
