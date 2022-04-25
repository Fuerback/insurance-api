package main

import (
	"log"
	"useorigin.com/insurance-api/internal/httpadapter/evaluationhttpadapter"
	"useorigin.com/insurance-api/internal/service/evaluationservice"
	"useorigin.com/insurance-api/server"
)

func main() {
	log.Println("Starting api server")

	service := evaluationservice.NewService()
	evaluationHandler := evaluationhttpadapter.NewEvaluationHandler(service)

	s := server.NewServer(evaluationHandler)
	s.Run()
}
