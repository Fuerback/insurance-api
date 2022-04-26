package insurancehttpadapter

import "useorigin.com/insurance-api/internal/service/insuranceservice"

type RiskProfileRequest struct {
	Age           *int     `json:"age" validate:"required,gte=0"`
	Dependents    *int     `json:"dependents" validate:"required,gte=0"`
	House         *House   `json:"house" validate:"omitempty"`
	Income        *int     `json:"income" validate:"required,gte=0"`
	MartialStatus string   `json:"marital_status" validate:"required,oneof=single married"`
	RiskQuestions []int    `json:"risk_questions" validate:"required,len=3,dive,gte=0,lte=1"`
	Vehicle       *Vehicle `json:"vehicle" validate:"omitempty"`
}

type House struct {
	OwnershipStatus string `json:"ownership_status" validate:"required,oneof=owned mortgaged"`
}

type Vehicle struct {
	Year int `json:"year" validate:"required,gte=0"`
}

func (u *RiskProfileRequest) toDomain() insuranceservice.RiskProfile {
	return insuranceservice.RiskProfile{
		Age:           *u.Age,
		Dependents:    *u.Dependents,
		House:         getDomainHouse(u.House),
		Income:        *u.Income,
		MartialStatus: u.MartialStatus,
		RiskQuestions: u.RiskQuestions,
		Vehicle:       getDomainVehicle(u.Vehicle),
	}
}

func getDomainVehicle(vehicle *Vehicle) *insuranceservice.Vehicle {
	if vehicle != nil {
		return &insuranceservice.Vehicle{
			Year: vehicle.Year,
		}
	}
	return nil
}

func getDomainHouse(house *House) *insuranceservice.House {
	if house != nil {
		return &insuranceservice.House{
			OwnershipStatus: house.OwnershipStatus,
		}
	}
	return nil
}

func NewEvaluation(age, dependents, income int, martialStatus string, riskQuestions []int, house *House, vehicle *Vehicle) RiskProfileRequest {
	return RiskProfileRequest{
		Age: func() *int {
			return &age
		}(),
		Dependents: func() *int {
			return &dependents
		}(),
		House: house,
		Income: func() *int {
			return &income
		}(),
		MartialStatus: martialStatus,
		RiskQuestions: riskQuestions,
		Vehicle:       vehicle,
	}
}
