package evaluationservice

type EvaluationService struct {
	rules []Rules
}

func NewService(rules []Rules) InsuranceEvaluation {
	return &EvaluationService{rules: rules}
}

func (e *EvaluationService) Evaluate(userInformation UserInformation) InsuranceSuggest {
	var initialRiskScore int8
	for _, answer := range userInformation.RiskQuestions {
		initialRiskScore += answer
	}

	riskScore := &InsuranceScore{
		Auto:       RiskScore{RiskPoint: initialRiskScore},
		Disability: RiskScore{RiskPoint: initialRiskScore},
		Home:       RiskScore{RiskPoint: initialRiskScore},
		Life:       RiskScore{RiskPoint: initialRiskScore},
	}

	for _, r := range e.rules {
		r.Evaluate(userInformation, riskScore)
	}

	return InsuranceSuggest{
		Auto:       getSuggest(riskScore.Auto),
		Disability: getSuggest(riskScore.Disability),
		Home:       getSuggest(riskScore.Home),
		Life:       getSuggest(riskScore.Life),
	}
}

func getSuggest(score RiskScore) string {
	if score.Ineligible == true {
		return "ineligible"
	}
	if score.RiskPoint <= 0 {
		return "economic"
	} else if score.RiskPoint >= 1 || score.RiskPoint <= 2 {
		return "regular"
	} else {
		return "responsible"
	}
}
