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

type insuranceProfile struct {
	RiskScore  int
	Ineligible bool
}

func NewInsuranceProfile(riskPoint int) insuranceProfile {
	return insuranceProfile{RiskScore: riskPoint}
}

func (rc *insuranceProfile) GetPlan() string {
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

type InsuranceScore struct {
	Auto       insuranceProfile
	Disability insuranceProfile
	Home       insuranceProfile
	Life       insuranceProfile
}
