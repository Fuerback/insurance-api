package main

import (
	"log"
	"useorigin.com/insurance-api/evaluation"
	"useorigin.com/insurance-api/server"
)

func main() {
	log.Println("Starting api server")

	// TODO: add logging and config

	evaluationHandler := evaluation.NewEvaluationHandler()

	s := server.NewServer(evaluationHandler)
	s.Run()
}
