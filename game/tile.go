package game

import (
	"github.com/beefsack/go-astar"
)

const (
	empty = iota
	food
	//headAround
	snake
)

const NoPassCost = 99999999999

var tileKindCost = map[int]float64{
	empty: 1.0,
	food:  1.0,
	//headAround: NoPassCost, // todo: change
	snake: NoPassCost,
}

type Tile struct {
	x, y            int
	board           *board
	costIndex       int
	snakeTileVanish int
}

func (t *Tile) Neighbors() []*Tile {
	var neighbors []*Tile
	for _, next := range NewMoves() {
		neighborTile, present := t.board.getTile(t.x+next.X, t.y+next.Y)
		if !present {
			continue
		}
		if tileKindCost[neighborTile.costIndex] < NoPassCost {
			neighbors = append(neighbors, neighborTile)
		}
	}
	return neighbors
}

// PathNeighbors repack into Pather interface
func (t *Tile) PathNeighbors() []astar.Pather {
	var neighbors []astar.Pather
	for _, n := range t.Neighbors() {
		neighbors = append(neighbors, n)
	}
	return neighbors
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	return to.(*Tile).Cost()
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	absX := toT.x - t.y
	if absX < 0 {
		absX = -absX
	}
	absY := toT.x - t.y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

func (t *Tile) Cost() float64 {
	return tileKindCost[t.costIndex]
}
