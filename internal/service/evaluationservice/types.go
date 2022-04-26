package evaluationservice

import "useorigin.com/insurance-api/internal/service/rules"

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

func (u *UserInformation) toRiskProfile(riskScore int) rules.RiskProfile {
	return rules.RiskProfile{
		Age:           u.Age,
		Dependents:    u.Dependents,
		House:         getRiskProfileHouse(u.House),
		Income:        u.Income,
		MartialStatus: u.MartialStatus,
		RiskQuestions: u.RiskQuestions,
		Vehicle:       getRiskProfileVehicle(u.Vehicle),
		RiskScore:     riskScore,
	}
}

func getRiskProfileVehicle(vehicle *Vehicle) *rules.Vehicle {
	if vehicle != nil {
		return &rules.Vehicle{
			Year: vehicle.Year,
		}
	}
	return nil
}

func getRiskProfileHouse(house *House) *rules.House {
	if house != nil {
		return &rules.House{
			OwnershipStatus: house.OwnershipStatus,
		}
	}
	return nil
}

type InsuranceSuggest struct {
	Auto       string `json:"auto"`
	Disability string `json:"disability"`
	Home       string `json:"home"`
	Life       string `json:"life"`
}
