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
	//
	//snake := []pairs{
	//	{0, 5}, {1, 5}, {2, 5}, {3, 5}, {4, 5}, {5, 5}, {5, 4}, {5, 3},
	//}
	//addSTestSnake(&g, snake)
	//snake = []pairs{
	//	//{1, 0}, {0, 0},
	//	{1, 2}, {1, 1}, {1, 0},
	//}
	//addTestYou(&g, snake)
	//food := []game.Coord{
	//	//{8, 8}, {7, 7},
	//	{2, 0}, {5, 5},
	//}
	//g.Board.Food = food
	////fmt.Println(s)
	//board := game.MakeBoard(g)
	//x, y, _ := game.FindFood(g.You.Head, board, g.Board.Food)
	////assert.Truef(t, ok, "nie ok")
	//best := game.FindCoordinates(x, y, g.You.Head)
	//fmt.Println("path ", best.Heading)
	//
	//moves := game.RankSpace(g.You.Head, board)
	//
	//t := time.Now()
	//game.Minimax(board, 15)
	//fmt.Printf("%v\n", time.Since(t))
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
