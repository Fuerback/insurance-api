package evaluationservice

import (
	"useorigin.com/insurance-api/internal/service/rules"
)

type EvaluationService struct{}

func NewService() InsuranceEvaluation {
	return &EvaluationService{}
}

func (e *EvaluationService) Evaluate(userInformation UserInformation) InsuranceSuggest {
	initialRiskScore := getInitialRiskScore(userInformation)

	// rule engine domain
	riskScore := &rules.InsuranceScore{
		Auto:       rules.NewInsuranceProfile(initialRiskScore),
		Disability: rules.NewInsuranceProfile(initialRiskScore),
		Home:       rules.NewInsuranceProfile(initialRiskScore),
		Life:       rules.NewInsuranceProfile(initialRiskScore),
	}

	// TODO: send 'rules1' and userProfile to GetInsuranceSuggest, it have to be responsible to Evaluate all rules and return an InsuranceSuggest (rules package can't have the service dependency)
	rules1 := loadRules()
	for _, r := range rules1 {
		r.Evaluate(userInformation, riskScore)
	}

	return InsuranceSuggest{
		Auto:       riskScore.Auto.GetPlan(),
		Disability: riskScore.Disability.GetPlan(),
		Home:       riskScore.Home.GetPlan(),
		Life:       riskScore.Life.GetPlan(),
	}
}

func loadRules() []rules.Rules {
	return []rules.Rules{
		rules.NewAutoRules(),
		rules.NewHomeRules(),
		rules.NewDisabilityRules(),
		rules.NewLifeRules(),
	}
}

func getInitialRiskScore(userInformation UserInformation) int {
	var initialRiskScore int
	for _, answer := range userInformation.RiskQuestions {
		initialRiskScore += answer
	}
	return initialRiskScore
}
