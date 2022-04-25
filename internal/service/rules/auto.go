package rules

import (
	"time"
	"useorigin.com/insurance-api/internal/service/evaluationservice"
)

type AutoRules struct{}

func NewAutoRules() Rules {
	return &AutoRules{}
}

func (r *AutoRules) Evaluate(userInformation evaluationservice.UserInformation, riskScore *InsuranceScore) {
	if userInformation.Vehicle == nil {
		riskScore.Auto.Ineligible = true
		return
	}
	if userInformation.Age < 30 {
		riskScore.Auto.RiskScore = riskScore.Auto.RiskScore - 2
	}
	if userInformation.Age >= 30 && userInformation.Age <= 40 {
		riskScore.Auto.RiskScore = riskScore.Auto.RiskScore - 1
	}
	if userInformation.Income > 200000 {
		riskScore.Auto.RiskScore = riskScore.Auto.RiskScore - 1
	}
	if userInformation.Vehicle != nil && userInformation.Vehicle.Year >= time.Now().Year()-5 {
		riskScore.Auto.RiskScore = riskScore.Auto.RiskScore + 1
	}
}
