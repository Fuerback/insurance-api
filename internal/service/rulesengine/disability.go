package rulesengine

type disabilityRules struct{}

func NewDisabilityRules() Rule {
	return &disabilityRules{}
}

func (r *disabilityRules) evaluate(riskProfile RiskProfile, profile *InsuranceProfile) {
	if riskProfile.Income == 0 {
		profile.Disability.Ineligible = true
		return
	}
	if riskProfile.Age > 60 {
		profile.Disability.Ineligible = true
		return
	}
	if riskProfile.Dependents > 0 {
		profile.Disability.RiskScore = profile.Disability.RiskScore + 1
	}
	if riskProfile.Age < 30 {
		profile.Disability.RiskScore = profile.Disability.RiskScore - 2
	}
	if riskProfile.Age >= 30 && riskProfile.Age <= 40 {
		profile.Disability.RiskScore = profile.Disability.RiskScore - 1
	}
	if riskProfile.Income > 200000 {
		profile.Disability.RiskScore = profile.Disability.RiskScore - 1
	}
	if riskProfile.House != nil && riskProfile.House.OwnershipStatus == MORTGAGED {
		profile.Disability.RiskScore = profile.Disability.RiskScore + 1
	}
	if riskProfile.MartialStatus == MARRIED {
		profile.Disability.RiskScore = profile.Disability.RiskScore - 1
	}
}
