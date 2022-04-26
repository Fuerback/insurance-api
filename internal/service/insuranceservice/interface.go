package insuranceservice

type Insurance interface {
	EvaluateUserProfile(userInformation RiskProfile) InsuranceSuggest
}
