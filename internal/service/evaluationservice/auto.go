package evaluationservice

import "time"

type AutoRules struct{}

func NewAutoRules() Rules {
	return &AutoRules{}
}

func (r *AutoRules) Evaluate(userInformation UserInformation, riskScore *InsuranceScore) {
	if userInformation.Vehicle == nil {
		riskScore.Auto.Ineligible = true
		return
	}
	if userInformation.Age < 30 {
		riskScore.Auto.RiskPoint = riskScore.Auto.RiskPoint - 2
	}
	if userInformation.Age >= 30 && userInformation.Age <= 40 {
		riskScore.Auto.RiskPoint = riskScore.Auto.RiskPoint - 1
	}
	if userInformation.Income > 200000 {
		riskScore.Auto.RiskPoint = riskScore.Auto.RiskPoint - 1
	}
	if userInformation.Vehicle != nil && userInformation.Vehicle.Year >= time.Now().Year()-5 {
		riskScore.Auto.RiskPoint = riskScore.Auto.RiskPoint + 1
	}
}
