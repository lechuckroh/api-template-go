package main

import (
	"github.com/lechuckroh/restapi-template-go/app"
	"log"
)

func main() {
	server := app.NewServer()
	if err := server.Run(8080); err != nil {
		log.Fatal("failed to run server", err)
	}
}
