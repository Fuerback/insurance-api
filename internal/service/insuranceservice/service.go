package insuranceservice

import (
	"useorigin.com/insurance-api/internal/service/rulesengine"
)

type InsuranceService struct{}

func NewService() Insurance {
	return &InsuranceService{}
}

func (e *InsuranceService) EvaluateUserProfile(riskProfile RiskProfile) InsuranceSuggest {
	initialRiskScore := getInitialRiskScore(riskProfile)

	evaluation := rulesengine.NewEvaluation(loadRules())
	profile := evaluation.Evaluate(riskProfile.toEngineRiskProfile(initialRiskScore))

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

func getInitialRiskScore(userInformation RiskProfile) int {
	var initialRiskScore int
	for _, answer := range userInformation.RiskQuestions {
		initialRiskScore += answer
	}
	return initialRiskScore
}
