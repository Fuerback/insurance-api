package main

import (
	"log"
	"useorigin.com/insurance-api/internal/httpadapter/insurancehttpadapter"
	"useorigin.com/insurance-api/internal/service/insuranceservice"
	"useorigin.com/insurance-api/server"
)

func main() {
	log.Println("Starting api server")

	service := insuranceservice.NewService()
	evaluationHandler := insurancehttpadapter.NewEvaluationHandler(service)

	s := server.NewServer(evaluationHandler)
	s.Run()
}
