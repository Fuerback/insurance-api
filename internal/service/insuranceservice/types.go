package insuranceservice

import "useorigin.com/insurance-api/internal/service/rulesengine"

type RiskProfile struct {
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

func (u *RiskProfile) toEngineRiskProfile(riskScore int) rulesengine.RiskProfile {
	return rulesengine.RiskProfile{
		Age:           u.Age,
		Dependents:    u.Dependents,
		House:         getRiskProfileHouse(u.House),
		Income:        u.Income,
		MartialStatus: u.MartialStatus,
		Vehicle:       getRiskProfileVehicle(u.Vehicle),
		RiskScore:     riskScore,
	}
}

func getRiskProfileVehicle(vehicle *Vehicle) *rulesengine.Vehicle {
	if vehicle != nil {
		return &rulesengine.Vehicle{
			Year: vehicle.Year,
		}
	}
	return nil
}

func getRiskProfileHouse(house *House) *rulesengine.House {
	if house != nil {
		return &rulesengine.House{
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
