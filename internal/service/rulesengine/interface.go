package rulesengine

type Rule interface {
	evaluate(riskProfile RiskProfile, profile *InsuranceProfile)
}

type Evaluation interface {
	Evaluate(riskProfile RiskProfile) InsuranceProfile
}
