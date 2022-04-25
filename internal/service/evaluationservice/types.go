package evaluationservice

type UserInformation struct {
	Age           int
	Dependents    int
	House         *House
	Income        int
	MartialStatus string
	RiskQuestions []int
	Vehicle       *Vehicle
}

type House struct {
	OwnershipStatus string
}

type Vehicle struct {
	Year int
}

type InsuranceSuggest struct {
	Auto       string `json:"auto"`
	Disability string `json:"disability"`
	Home       string `json:"home"`
	Life       string `json:"life"`
}

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

type RiskScore struct {
	RiskPoint  int
	Ineligible bool
}

func NewRiskScore(riskPoint int) RiskScore {
	return RiskScore{RiskPoint: riskPoint}
}

func (rc *RiskScore) GetPlan() string {
	if rc.Ineligible == true {
		return INELIGIBLE
	} else if rc.RiskPoint <= 0 {
		return ECONOMIC
	} else if rc.RiskPoint >= 1 && rc.RiskPoint <= 2 {
		return REGULAR
	} else {
		return RESPONSIBLE
	}
}

type InsuranceScore struct {
	Auto       RiskScore
	Disability RiskScore
	Home       RiskScore
	Life       RiskScore
}
