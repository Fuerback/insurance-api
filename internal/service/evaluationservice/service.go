package evaluationservice

import (
	"useorigin.com/insurance-api/internal/service/rulesengine"
)

type EvaluationService struct{}

func NewService() InsuranceEvaluation {
	return &EvaluationService{}
}

func (e *EvaluationService) Evaluate(userInformation UserInformation) InsuranceSuggest {
	initialRiskScore := getInitialRiskScore(userInformation)

	evaluation := rulesengine.NewEvaluation(loadRules())
	profile := evaluation.Evaluate(userInformation.toRiskProfile(initialRiskScore))

	return InsuranceSuggest{
		Auto:       profile.Auto.GetPlan(),
		Disability: profile.Disability.GetPlan(),
		Home:       profile.Home.GetPlan(),
		Life:       profile.Life.GetPlan(),
	}
}

func loadRules() []rulesengine.Rule {
	return []rulesengine.Rule{
		rulesengine.NewAutoRules(),
		rulesengine.NewHomeRules(),
		rulesengine.NewDisabilityRules(),
		rulesengine.NewLifeRules(),
	}
}

func getInitialRiskScore(userInformation UserInformation) int {
	var initialRiskScore int
	for _, answer := range userInformation.RiskQuestions {
		initialRiskScore += answer
	}
	return initialRiskScore
}
