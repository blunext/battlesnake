package models

import (
	"fmt"
	"github.com/beefsack/go-astar"
)

const (
	empty = iota
	food
	//headAround
	snake
	wall
)

const NoPassCost = 99999999999

var tileKindCost = map[int]float64{
	empty: 1.0,
	food:  1.0,
	//headAround: NoPassCost, // todo: change
	snake: NoPassCost,
	wall:  NoPassCost,
}

type Tile struct {
	X, Y        int
	board       *MyBoard
	costIndex   int
	snakeTileNo int
}

func (t *Tile) Neighbors() []*Tile {
	var neighbors []*Tile
	for _, next := range NewMoves() {
		if t.X == 1 && t.Y == 2 {
			fmt.Println()
		}
		neighborTile := t.board.Tile(t.X+next.X, t.Y+next.Y)
		if neighborTile.snakeTileNo == 1 {
			continue
		}
		neighbors = append(neighbors, neighborTile)
	}
	return neighbors
}

// PathNeighbors repack into Pather interface
func (t *Tile) PathNeighbors() []astar.Pather {
	var neighbors []astar.Pather
	ns := t.Neighbors()
	for _, n := range ns {
		if tileKindCost[n.costIndex] < NoPassCost {
			neighbors = append(neighbors, n)
		}
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
