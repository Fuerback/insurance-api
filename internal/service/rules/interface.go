package rules

import "useorigin.com/insurance-api/internal/service/evaluationservice"

type Rules interface {
	// TODO: remove evaluationservice dependency
	Evaluate(userInformation evaluationservice.UserInformation, riskScore *InsuranceScore)
}
