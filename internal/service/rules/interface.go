package rules

type Rules interface {
	Evaluate(riskScore *RiskScore)
}
