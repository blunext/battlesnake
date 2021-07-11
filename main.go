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
