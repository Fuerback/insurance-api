package evaluationservice

const (
	MORTGAGED = "mortgaged"
	MARRIED   = "married"
)

type UserInformation struct {
	Age           int
	Dependents    int
	House         *House
	Income        int
	MartialStatus string
	RiskQuestions []int8
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

type RiskScore struct {
	RiskPoint  int8
	Ineligible bool
}

type InsuranceScore struct {
	Auto       RiskScore
	Disability RiskScore
	Home       RiskScore
	Life       RiskScore
}
