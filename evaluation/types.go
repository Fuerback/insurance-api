package evaluation

type Evaluation struct {
	Age           int     `json:"age" validate:"required,gte=0"`
	Dependents    int     `json:"dependents" validate:"required,gte=0"`
	House         House   `json:"house" validate:"dive"`
	Income        int     `json:"income" validate:"required,gte=0"`
	MartialStatus string  `json:"marital_status" validate:"required"`
	RiskQuestions []bool  `json:"risk_questions" validate:"required,len=3"`
	Vehicle       Vehicle `json:"vehicle" validate:"dive"`
}

type House struct {
	OwnershipStatus string `json:"ownership_status" validate:"required"`
}

type Vehicle struct {
	Year int `json:"year" validate:"required,gte=0"`
}

type Error struct {
	Message string `json:"message"`
}
