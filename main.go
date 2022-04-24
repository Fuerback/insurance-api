package main

import (
	"log"
	"useorigin.com/insurance-api/internal/server"
)

func main() {
	log.Println("Starting api server")

	// TODO: add logging and config

	s := server.NewServer()
	s.Run()
}
