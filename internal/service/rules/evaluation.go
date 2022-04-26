package rules

type evaluation struct {
	rules []Rule
}

func NewEvaluation(rules []Rule) *evaluation {
	return &evaluation{rules: rules}
}

func (e *evaluation) Evaluate(riskProfile RiskProfile) InsuranceProfile {
	insuranceScore := NewRiskScore(riskProfile.RiskScore)
	profile := NewInsuranceProfile(insuranceScore)
	for _, r := range e.rules {
		r.evaluate(riskProfile, profile)
	}
	return *profile
}
