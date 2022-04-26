package rulesengine

type Rule interface {
	evaluate(riskProfile RiskProfile, profile *InsuranceProfile)
}

type Evaluation interface {
	EvaluateRules(riskProfile RiskProfile) InsuranceProfile
}
