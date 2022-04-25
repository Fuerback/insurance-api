package evaluation

type Evaluation struct {
	Age           *int     `json:"age" validate:"required,gte=0"`
	Dependents    *int     `json:"dependents" validate:"required,gte=0"`
	House         *House   `json:"house" validate:"omitempty"`
	Income        *int     `json:"income" validate:"required,gte=0"`
	MartialStatus string   `json:"marital_status" validate:"required,oneof=single married"`
	RiskQuestions []int    `json:"risk_questions" validate:"required,len=3"`
	Vehicle       *Vehicle `json:"vehicle" validate:"omitempty"`
}

type House struct {
	OwnershipStatus string `json:"ownership_status" validate:"required,oneof=owned mortgaged"`
}

type Vehicle struct {
	Year *int `json:"year" validate:"required,gte=0"`
}
