package rules

const (
	INELIGIBLE  = "ineligible"
	ECONOMIC    = "economic"
	REGULAR     = "regular"
	RESPONSIBLE = "responsible"
	MORTGAGED   = "mortgaged"
	MARRIED     = "married"
	SINGLE      = "single"
	OWNED       = "owned"
)

type RiskProfile struct {
	Age           int
	Dependents    int
	House         *House
	Income        int
	MartialStatus string
	RiskQuestions []int
	Vehicle       *Vehicle
	RiskScore     int
}

type House struct {
	OwnershipStatus string
}

type Vehicle struct {
	Year int
}

type riskScore struct {
	RiskScore  int
	Ineligible bool
}

func NewRiskScore(score int) riskScore {
	return riskScore{RiskScore: score}
}

func (rc *riskScore) GetPlan() string {
	if rc.Ineligible == true {
		return INELIGIBLE
	} else if rc.RiskScore <= 0 {
		return ECONOMIC
	} else if rc.RiskScore >= 1 && rc.RiskScore <= 2 {
		return REGULAR
	} else {
		return RESPONSIBLE
	}
}

type InsuranceProfile struct {
	Auto       riskScore
	Disability riskScore
	Home       riskScore
	Life       riskScore
}

func NewInsuranceProfile(score riskScore) *InsuranceProfile {
	return &InsuranceProfile{Auto: score, Disability: score, Home: score, Life: score}
}
