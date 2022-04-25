package evaluationservice

type EvaluationService struct {
	rules []Rules
}

func NewService(rules []Rules) InsuranceEvaluation {
	return &EvaluationService{rules: rules}
}

func (e *EvaluationService) Evaluate(userInformation UserInformation) InsuranceSuggest {
	var initialRiskScore int
	for _, answer := range userInformation.RiskQuestions {
		initialRiskScore += answer
	}

	// rule engine domain
	riskScore := &InsuranceScore{
		Auto:       NewRiskScore(initialRiskScore),
		Disability: NewRiskScore(initialRiskScore),
		Home:       NewRiskScore(initialRiskScore),
		Life:       NewRiskScore(initialRiskScore),
	}

	// create a userInformation to rule engine?

	for _, r := range e.rules {
		r.Evaluate(userInformation, riskScore)
	}

	return InsuranceSuggest{
		Auto:       riskScore.Auto.GetPlan(),
		Disability: riskScore.Disability.GetPlan(),
		Home:       riskScore.Home.GetPlan(),
		Life:       riskScore.Life.GetPlan(),
	}
}
