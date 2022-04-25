package rules

import "useorigin.com/insurance-api/internal/service/evaluationservice"

type LifeRules struct{}

func NewLifeRules() Rules {
	return &LifeRules{}
}

func (r *LifeRules) Evaluate(userInformation evaluationservice.UserInformation, riskScore *InsuranceScore) {
	if userInformation.Age > 60 {
		riskScore.Life.Ineligible = true
		return
	}
	if userInformation.Dependents > 0 {
		riskScore.Life.RiskScore = riskScore.Life.RiskScore + 1
	}
	if userInformation.Age < 30 {
		riskScore.Life.RiskScore = riskScore.Life.RiskScore - 2
	}
	if userInformation.Age >= 30 && userInformation.Age <= 40 {
		riskScore.Life.RiskScore = riskScore.Life.RiskScore - 1
	}
	if userInformation.Income > 200000 {
		riskScore.Life.RiskScore = riskScore.Life.RiskScore - 1
	}
	if userInformation.MartialStatus == MARRIED {
		riskScore.Life.RiskScore = riskScore.Life.RiskScore + 1
	}
}
