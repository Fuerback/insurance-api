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
