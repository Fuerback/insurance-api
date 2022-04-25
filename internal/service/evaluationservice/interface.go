package evaluationservice

type InsuranceEvaluation interface {
	Evaluate(userInformation UserInformation) InsuranceSuggest
}
