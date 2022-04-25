package main

import (
	"log"
	"useorigin.com/insurance-api/internal/httpadapter/evaluationhttpadapter"
	"useorigin.com/insurance-api/internal/service/evaluationservice"
	"useorigin.com/insurance-api/server"
)

func main() {
	log.Println("Starting api server")

	// TODO: add logging and config

	rules := make([]evaluationservice.Rules, 0)
	rules = append(rules, evaluationservice.NewAutoRules())
	rules = append(rules, evaluationservice.NewHomeRules())
	rules = append(rules, evaluationservice.NewDisabilityRules())
	rules = append(rules, evaluationservice.NewLifeRules())

	service := evaluationservice.NewService(rules)
	evaluationHandler := evaluationhttpadapter.NewEvaluationHandler(service)

	s := server.NewServer(evaluationHandler)
	s.Run()
}
