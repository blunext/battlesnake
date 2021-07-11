package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"snakehero/web"
)

func main() {
	//
	////moves := newMoves()
	//g := game.GameRequest{}
	//g.Board.Width = 11
	//g.Board.Height = 11
	////g.You.Head.X = 5
	////g.You.Head.Y = 5
	//
	//snake := []pairs{
	//	{0, 3}, {1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3}, {5, 2}, {5, 1},
	//}
	//addSnake(&g, snake)
	//snake = []pairs{
	//	{1, 0}, {0, 0},
	//}
	//addYou(&g, snake)
	//food := []game.Coord{
	//	{8, 8}, {7, 6},
	//}
	//g.Board.Food = food
	////fmt.Println(s)
	//board := game.MakeBoard(g)
	//
	//game.FindFood(g.You.Head, board, g.Board.Food)
	//moves := game.RankSpace(g.You.Head, board)
	//
	//fmt.Println(game.FindBest(moves))

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", web.HandleIndex)
	http.HandleFunc("/start", web.HandleStart)
	http.HandleFunc("/move", web.HandleMove)
	http.HandleFunc("/end", web.HandleEnd)

	fmt.Printf("Starting Battlesnake Server at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
