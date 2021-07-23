package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"snakehero/game"
	"snakehero/web"
	"testing"
	"time"
)

func TestAll(t *testing.T) {

	//moves := newMoves()
	g := game.GameRequest{}
	g.Board.Width = 11
	g.Board.Height = 11

	snake := [][]int{
		{0, 5}, {1, 5}, {2, 5}, {3, 5}, {4, 5}, {5, 5}, {5, 4}, {5, 3},
	}
	game.AddSTestSnake(&g, snake)

	//snake = [][]int{
	//	{0, 8}, {1, 8}, {2, 8}, {3, 8}, {4, 8}, {5, 8},
	//}
	//game.AddSTestSnake(&g, snake)

	snake = [][]int{
		//{1, 0}, {0, 0},
		{1, 2}, {1, 1}, {1, 0},
	}
	game.AddTestYou(&g, snake)
	food := []game.Coord{
		//{8, 8}, {7, 7},
		{2, 2}, {5, 5},
	}
	g.Board.Food = food
	//fmt.Println(s)
	board := game.MakeBoard(g)
	x, y, ok := game.FindFood(g.You.Head, board, g.Board.Food)
	assert.Truef(t, ok, "nie ok")
	best := game.FindCoordinates(x, y, g.You.Head)
	fmt.Println("path ", best.Heading)

	//moves := game.RankSpace(g.You.Head, board)

	tm := time.Now()
	round := game.Minimax(board, game.MMdepth, board.GameData.You.ID)
	fmt.Printf("counter %v\n", game.Counter)
	for _, r := range round {
		if r.SnakeId == board.GameData.You.ID {
			fmt.Printf("move: %v\n", r.Move)
		}
	}
	fmt.Printf("%v\n", time.Since(tm))
	//
	//fmt.Println(game.FindBest(moves))

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	a := `{"game":{"id":"dd738c2b-e863-4d94-be16-eda3e40f36ef","timeout":500},"turn":102,"board":{"height":11,"width":11,"food":[{"x":5,"y":5},{"x":2,"y":5},{"x":0,"y":4},{"x":3,"y":6},{"x":2,"y":8},{"x":4,"y":10},{"x":0,"y":6},{"x":0,"y":7},{"x":2,"y":0},{"x":10,"y":0},{"x":5,"y":3},{"x":8,"y":0},{"x":7,"y":4},{"x":6,"y":3},{"x":0,"y":10}],"hazards":[],"snakes":[{"id":"66f0e089-5cf6-484f-8fda-b541f4b75769","name":"Snake1","health":97,"body":[{"x":9,"y":10},{"x":10,"y":10},{"x":10,"y":9},{"x":10,"y":8},{"x":10,"y":7},{"x":9,"y":7},{"x":9,"y":8},{"x":8,"y":8}],"latency":0,"head":{"x":9,"y":10},"length":8,"shout":"","squad":""}]},"you":{"id":"66f0e089-5cf6-484f-8fda-b541f4b75769","name":"Snake1","health":97,"body":[{"x":9,"y":10},{"x":10,"y":10},{"x":10,"y":9},{"x":10,"y":8},{"x":10,"y":7},{"x":9,"y":7},{"x":9,"y":8},{"x":8,"y":8}],"latency":0,"head":{"x":9,"y":10},"length":8,"shout":"","squad":""}}`
	code, _ := makePost(t, []byte(a), web.HandleMove())
	assert.Equal(t, 200, code, "response code not expected")

	// read
}

func TestFromJson(t *testing.T) {
	a := `{"game":{"id":"dd738c2b-e863-4d94-be16-eda3e40f36ef","timeout":500},"turn":102,"board":{"height":11,"width":11,"food":[{"x":5,"y":5},{"x":2,"y":5},{"x":0,"y":4},{"x":3,"y":6},{"x":2,"y":8},{"x":4,"y":10},{"x":0,"y":6},{"x":0,"y":7},{"x":2,"y":0},{"x":10,"y":0},{"x":5,"y":3},{"x":8,"y":0},{"x":7,"y":4},{"x":6,"y":3},{"x":0,"y":10}],"hazards":[],"snakes":[{"id":"66f0e089-5cf6-484f-8fda-b541f4b75769","name":"Snake1","health":97,"body":[{"x":9,"y":10},{"x":10,"y":10},{"x":10,"y":9},{"x":10,"y":8},{"x":10,"y":7},{"x":9,"y":7},{"x":9,"y":8},{"x":8,"y":8}],"latency":0,"head":{"x":9,"y":10},"length":8,"shout":"","squad":""}]},"you":{"id":"66f0e089-5cf6-484f-8fda-b541f4b75769","name":"Snake1","health":97,"body":[{"x":9,"y":10},{"x":10,"y":10},{"x":10,"y":9},{"x":10,"y":8},{"x":10,"y":7},{"x":9,"y":7},{"x":9,"y":8},{"x":8,"y":8}],"latency":0,"head":{"x":9,"y":10},"length":8,"shout":"","squad":""}}`
	code, _ := makePost(t, []byte(a), web.HandleMove())
	assert.Equal(t, 200, code, "response code not expected")

	a = `{"game":{"id":"8399b015-49e2-4a1b-ae87-1806bf61ce60","timeout":500},"turn":183,"board":{"height":11,"width":11,"food":[{"x":4,"y":2},{"x":4,"y":3},{"x":0,"y":4},{"x":0,"y":1},{"x":5,"y":3},{"x":10,"y":0},{"x":7,"y":1},{"x":10,"y":1},{"x":5,"y":2},{"x":0,"y":0},{"x":4,"y":7},{"x":1,"y":0},{"x":1,"y":1},{"x":5,"y":10},{"x":4,"y":1},{"x":10,"y":2},{"x":9,"y":0},{"x":3,"y":7},{"x":7,"y":2},{"x":5,"y":0},{"x":10,"y":6},{"x":6,"y":10},{"x":8,"y":3},{"x":3,"y":3},{"x":7,"y":10},{"x":2,"y":10},{"x":9,"y":3}],"hazards":[],"snakes":[{"id":"4f9cfca9-e37a-484c-a05e-40011192624f","name":"Snake1","health":95,"body":[{"x":0,"y":8},{"x":0,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":1,"y":7},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":1,"y":5},{"x":2,"y":5},{"x":2,"y":4},{"x":3,"y":4}],"latency":0,"head":{"x":0,"y":8},"length":12,"shout":"","squad":""}]},"you":{"id":"4f9cfca9-e37a-484c-a05e-40011192624f","name":"Snake1","health":95,"body":[{"x":0,"y":8},{"x":0,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":1,"y":7},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":1,"y":5},{"x":2,"y":5},{"x":2,"y":4},{"x":3,"y":4}],"latency":0,"head":{"x":0,"y":8},"length":12,"shout":"","squad":""}}`
	code, _ = makePost(t, []byte(a), web.HandleMove())
	assert.Equal(t, 200, code, "response code not expected")

}

func makePost(t *testing.T, jsonMessage []byte, handler http.Handler) (int, string) {
	req, err := http.NewRequest("POST", "/move", bytes.NewBuffer(jsonMessage))
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	handler.ServeHTTP(r, req)
	return r.Code, r.Body.String()
}

func makePostB(t *testing.B, jsonMessage []byte, handler http.Handler) (int, string) {
	req, err := http.NewRequest("POST", "/move", bytes.NewBuffer(jsonMessage))
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	handler.ServeHTTP(r, req)
	return r.Code, r.Body.String()
}
func jsonFromStruct(s interface{}) []byte {
	j, err := json.Marshal(s)
	if err != nil {
		log.Panicf("cannot marshal %v, err: %v", s, err)
	}
	return j
}
