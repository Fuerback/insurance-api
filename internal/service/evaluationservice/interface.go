package evaluationservice

type InsuranceEvaluation interface {
	Evaluate(userInformation UserInformation) InsuranceSuggest
}

type Rules interface {
	Evaluate(userInformation UserInformation, riskScore *InsuranceScore)
}
