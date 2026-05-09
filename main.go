package main

import (
	"log"

	"github.com/misiuleczek02/zadanie7-server/internal/server"
)

func main() {
	app := server.New()
	if err := app.Start(":8080"); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
