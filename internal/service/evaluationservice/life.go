package evaluationservice

type LifeRules struct{}

func NewLifeRules() Rules {
	return &LifeRules{}
}

func (r *LifeRules) Evaluate(userInformation UserInformation, riskScore *InsuranceScore) {
	if userInformation.Age > 60 {
		riskScore.Life.Ineligible = true
		return
	}
	if userInformation.Dependents > 0 {
		riskScore.Life.RiskPoint = riskScore.Life.RiskPoint + 1
	}
	if userInformation.Age < 30 {
		riskScore.Life.RiskPoint = riskScore.Life.RiskPoint - 2
	}
	if userInformation.Age >= 30 && userInformation.Age <= 40 {
		riskScore.Life.RiskPoint = riskScore.Life.RiskPoint - 1
	}
	if userInformation.Income > 200000 {
		riskScore.Life.RiskPoint = riskScore.Life.RiskPoint - 1
	}
	if userInformation.MartialStatus == MARRIED {
		riskScore.Life.RiskPoint = riskScore.Life.RiskPoint + 1
	}
}
