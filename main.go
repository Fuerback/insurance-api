package main

import (
	"log"
	"useorigin.com/insurance-api/internal/httpadapter/evaluationhttpadapter"
	"useorigin.com/insurance-api/server"
)

func main() {
	log.Println("Starting api server")

	// TODO: add logging and config

	evaluationHandler := evaluationhttpadapter.NewEvaluationHandler()

	s := server.NewServer(evaluationHandler)
	s.Run()
}
