package insuranceservice

import (
	"useorigin.com/insurance-api/internal/service/rulesengine"
)

type InsuranceService struct {
	evaluation rulesengine.Evaluation
}

func NewService(evaluation rulesengine.Evaluation) Insurance {
	return &InsuranceService{evaluation: evaluation}
}

func (e *InsuranceService) EvaluateUserProfile(riskProfile RiskProfile) InsuranceSuggest {
	initialRiskScore := getInitialRiskScore(riskProfile)

	profile := e.evaluation.EvaluateRules(riskProfile.toEngineRiskProfile(initialRiskScore))

	return InsuranceSuggest{
		Auto:       profile.Auto.GetPlan(),
		Disability: profile.Disability.GetPlan(),
		Home:       profile.Home.GetPlan(),
		Life:       profile.Life.GetPlan(),
	}
}

func getInitialRiskScore(userInformation RiskProfile) int {
	var initialRiskScore int
	for _, answer := range userInformation.RiskQuestions {
		initialRiskScore += answer
	}
	return initialRiskScore
}
