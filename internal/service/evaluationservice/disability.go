package evaluationservice

type DisabilityRules struct{}

func NewDisabilityRules() Rules {
	return &DisabilityRules{}
}

func (r *DisabilityRules) Evaluate(userInformation UserInformation, riskScore *InsuranceScore) {
	if userInformation.Income == 0 {
		riskScore.Disability.Ineligible = true
		return
	}
	if userInformation.Age > 60 {
		riskScore.Disability.Ineligible = true
		return
	}
	if userInformation.Dependents > 0 {
		riskScore.Disability.RiskPoint = riskScore.Disability.RiskPoint + 1
	}
	if userInformation.Age < 30 {
		riskScore.Disability.RiskPoint = riskScore.Disability.RiskPoint - 2
	}
	if userInformation.Age >= 30 && userInformation.Age <= 40 {
		riskScore.Disability.RiskPoint = riskScore.Disability.RiskPoint - 1
	}
	if userInformation.Income > 200000 {
		riskScore.Disability.RiskPoint = riskScore.Disability.RiskPoint - 1
	}
	if userInformation.House != nil && userInformation.House.OwnershipStatus == MORTGAGED {
		riskScore.Disability.RiskPoint = riskScore.Disability.RiskPoint + 1
	}
	if userInformation.MartialStatus == MARRIED {
		riskScore.Disability.RiskPoint = riskScore.Disability.RiskPoint - 1
	}
}
