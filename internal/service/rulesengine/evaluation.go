package rulesengine

type engine struct {
	rules []Rule
}

func NewEngine() *engine {
	return &engine{rules: loadRules()}
}

func (e *engine) EvaluateRules(riskProfile RiskProfile) InsuranceProfile {
	insuranceScore := NewRiskScore(riskProfile.RiskScore)
	profile := NewInsuranceProfile(insuranceScore)
	for _, r := range e.rules {
		r.evaluate(riskProfile, profile)
	}
	return *profile
}

func loadRules() []Rule {
	return []Rule{
		newAutoRules(),
		newHomeRules(),
		newDisabilityRules(),
		newLifeRules(),
	}
}
