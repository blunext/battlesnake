package models

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
	X, Y            int
	board           *MyBoard
	costIndex       int
	snakeTileVanish int
}

func (t *Tile) Neighbors() []*Tile {
	var neighbors []*Tile
	for _, next := range NewMoves() {
		neighborTile, present := t.board.GetTile(t.X+next.X, t.Y+next.Y)
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
	absX := toT.X - t.Y
	if absX < 0 {
		absX = -absX
	}
	absY := toT.X - t.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

func (t *Tile) Cost() float64 {
	return tileKindCost[t.costIndex]
}

func NewMoves() []Direction {
	return []Direction{
		{0, 1, "up", 0},
		{0, -1, "down", 0},
		{-1, 0, "left", 0},
		{1, 0, "right", 0},
	}
}
