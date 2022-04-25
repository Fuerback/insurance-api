package evaluationservice

type HomeRules struct{}

func NewHomeRules() Rules {
	return &HomeRules{}
}

func (r *HomeRules) Evaluate(userInformation UserInformation, riskScore *InsuranceScore) {
	if userInformation.House == nil {
		riskScore.Home.Ineligible = true
		return
	}
	if userInformation.Age < 30 {
		riskScore.Home.RiskPoint = riskScore.Home.RiskPoint - 2
	}
	if userInformation.Age >= 30 && userInformation.Age <= 40 {
		riskScore.Home.RiskPoint = riskScore.Home.RiskPoint - 1
	}
	if userInformation.Income > 200000 {
		riskScore.Home.RiskPoint = riskScore.Home.RiskPoint - 1
	}
	if userInformation.House != nil && userInformation.House.OwnershipStatus == MORTGAGED {
		riskScore.Home.RiskPoint = riskScore.Home.RiskPoint + 1
	}
}
