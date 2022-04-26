package rulesengine

import (
	"time"
)

type autoRules struct{}

func newAutoRules() Rule {
	return &autoRules{}
}

func (r *autoRules) evaluate(riskProfile RiskProfile, profile *InsuranceProfile) {
	if riskProfile.Vehicle == nil {
		profile.Auto.Ineligible = true
		return
	}
	if riskProfile.Age < 30 {
		profile.Auto.RiskScore = profile.Auto.RiskScore - 2
	}
	if riskProfile.Age >= 30 && riskProfile.Age <= 40 {
		profile.Auto.RiskScore = profile.Auto.RiskScore - 1
	}
	if riskProfile.Income > 200000 {
		profile.Auto.RiskScore = profile.Auto.RiskScore - 1
	}
	if riskProfile.Vehicle != nil && riskProfile.Vehicle.Year >= time.Now().Year()-5 {
		profile.Auto.RiskScore = profile.Auto.RiskScore + 1
	}
}
