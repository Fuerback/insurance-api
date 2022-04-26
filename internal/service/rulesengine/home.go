package rulesengine

type homeRules struct{}

func NewHomeRules() Rule {
	return &homeRules{}
}

func (r *homeRules) evaluate(riskProfile RiskProfile, profile *InsuranceProfile) {
	if riskProfile.House == nil {
		profile.Home.Ineligible = true
		return
	}
	if riskProfile.Age < 30 {
		profile.Home.RiskScore = profile.Home.RiskScore - 2
	}
	if riskProfile.Age >= 30 && riskProfile.Age <= 40 {
		profile.Home.RiskScore = profile.Home.RiskScore - 1
	}
	if riskProfile.Income > 200000 {
		profile.Home.RiskScore = profile.Home.RiskScore - 1
	}
	if riskProfile.House != nil && riskProfile.House.OwnershipStatus == MORTGAGED {
		profile.Home.RiskScore = profile.Home.RiskScore + 1
	}
}
