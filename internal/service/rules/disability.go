package rules

import "useorigin.com/insurance-api/internal/service/evaluationservice"

type DisabilityRules struct{}

func NewDisabilityRules() Rules {
	return &DisabilityRules{}
}

func (r *DisabilityRules) Evaluate(userInformation evaluationservice.UserInformation, riskScore *InsuranceScore) {
	if userInformation.Income == 0 {
		riskScore.Disability.Ineligible = true
		return
	}
	if userInformation.Age > 60 {
		riskScore.Disability.Ineligible = true
		return
	}
	if userInformation.Dependents > 0 {
		riskScore.Disability.RiskScore = riskScore.Disability.RiskScore + 1
	}
	if userInformation.Age < 30 {
		riskScore.Disability.RiskScore = riskScore.Disability.RiskScore - 2
	}
	if userInformation.Age >= 30 && userInformation.Age <= 40 {
		riskScore.Disability.RiskScore = riskScore.Disability.RiskScore - 1
	}
	if userInformation.Income > 200000 {
		riskScore.Disability.RiskScore = riskScore.Disability.RiskScore - 1
	}
	if userInformation.House != nil && userInformation.House.OwnershipStatus == MORTGAGED {
		riskScore.Disability.RiskScore = riskScore.Disability.RiskScore + 1
	}
	if userInformation.MartialStatus == MARRIED {
		riskScore.Disability.RiskScore = riskScore.Disability.RiskScore - 1
	}
}
