package rulesengine

type lifeRules struct{}

func newLifeRules() Rule {
	return &lifeRules{}
}

func (r *lifeRules) evaluate(riskProfile RiskProfile, profile *InsuranceProfile) {
	if riskProfile.Age > 60 {
		profile.Life.Ineligible = true
		return
	}
	if riskProfile.Dependents > 0 {
		profile.Life.RiskScore = profile.Life.RiskScore + 1
	}
	if riskProfile.Age < 30 {
		profile.Life.RiskScore = profile.Life.RiskScore - 2
	}
	if riskProfile.Age >= 30 && riskProfile.Age <= 40 {
		profile.Life.RiskScore = profile.Life.RiskScore - 1
	}
	if riskProfile.Income > 200000 {
		profile.Life.RiskScore = profile.Life.RiskScore - 1
	}
	if riskProfile.MartialStatus == MARRIED {
		profile.Life.RiskScore = profile.Life.RiskScore + 1
	}
}
