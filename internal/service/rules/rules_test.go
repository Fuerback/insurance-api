package rules

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAutoRules(t *testing.T) {
	var tests = []struct {
		riskProfile RiskProfile
		plan        string
	}{
		{RiskProfile{
			Age:        10,
			Dependents: 1,
			House: func() *House {
				house := House{OwnershipStatus: OWNED}
				return &house
			}(),
			Income:        0,
			MartialStatus: MARRIED,
			Vehicle: func() *Vehicle {
				vehicle := Vehicle{Year: 2010}
				return &vehicle
			}(),
			RiskScore: 2,
		},
			ECONOMIC,
		},
		{RiskProfile{
			Age:        35,
			Dependents: 1,
			House: func() *House {
				house := House{OwnershipStatus: OWNED}
				return &house
			}(),
			Income:        0,
			MartialStatus: MARRIED,
			Vehicle: func() *Vehicle {
				vehicle := Vehicle{Year: 2010}
				return &vehicle
			}(),
			RiskScore: 3,
		},
			REGULAR,
		},
		{RiskProfile{
			Age:        35,
			Dependents: 1,
			House: func() *House {
				house := House{OwnershipStatus: OWNED}
				return &house
			}(),
			Income:        0,
			MartialStatus: MARRIED,
			RiskScore:     3,
		},
			INELIGIBLE,
		},
		{RiskProfile{
			Age:        65,
			Dependents: 1,
			House: func() *House {
				house := House{OwnershipStatus: OWNED}
				return &house
			}(),
			Income:        1000,
			MartialStatus: MARRIED,
			Vehicle: func() *Vehicle {
				vehicle := Vehicle{Year: 2020}
				return &vehicle
			}(),
			RiskScore: 2,
		},
			RESPONSIBLE,
		},
		{RiskProfile{
			Age:        65,
			Dependents: 1,
			House: func() *House {
				house := House{OwnershipStatus: OWNED}
				return &house
			}(),
			Income:        250000,
			MartialStatus: MARRIED,
			Vehicle: func() *Vehicle {
				vehicle := Vehicle{Year: 2020}
				return &vehicle
			}(),
			RiskScore: 2,
		},
			REGULAR,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.plan)
		t.Run(testname, func(t *testing.T) {
			insuranceScore := NewRiskScore(tt.riskProfile.RiskScore)
			profile := NewInsuranceProfile(insuranceScore)
			autoRules := NewAutoRules()
			autoRules.evaluate(tt.riskProfile, profile)

			assert.Equal(t, tt.plan, profile.Auto.GetPlan())
		})
	}
}

func TestDisabilityRules(t *testing.T) {
	var tests = []struct {
		riskProfile RiskProfile
		plan        string
	}{
		{RiskProfile{
			Age:        10,
			Dependents: 1,
			House: func() *House {
				house := House{OwnershipStatus: OWNED}
				return &house
			}(),
			Income:        0,
			MartialStatus: MARRIED,
			Vehicle: func() *Vehicle {
				vehicle := Vehicle{Year: 2010}
				return &vehicle
			}(),
			RiskScore: 2,
		},
			INELIGIBLE,
		},
		{RiskProfile{
			Age:        61,
			Dependents: 1,
			House: func() *House {
				house := House{OwnershipStatus: OWNED}
				return &house
			}(),
			Income:        10,
			MartialStatus: MARRIED,
			Vehicle: func() *Vehicle {
				vehicle := Vehicle{Year: 2010}
				return &vehicle
			}(),
			RiskScore: 1,
		},
			INELIGIBLE,
		},
		{RiskProfile{
			Age:        29,
			Dependents: 1,
			House: func() *House {
				house := House{OwnershipStatus: MORTGAGED}
				return &house
			}(),
			Income:        10000,
			MartialStatus: MARRIED,
			RiskScore:     0,
		},
			ECONOMIC,
		},
		{RiskProfile{
			Age:        35,
			Dependents: 1,
			House: func() *House {
				house := House{OwnershipStatus: OWNED}
				return &house
			}(),
			Income:        200001,
			MartialStatus: SINGLE,
			Vehicle: func() *Vehicle {
				vehicle := Vehicle{Year: 2010}
				return &vehicle
			}(),
			RiskScore: 1,
		},
			ECONOMIC,
		},
		{RiskProfile{
			Age:           59,
			Dependents:    1,
			Income:        50000,
			MartialStatus: MARRIED,
			RiskScore:     3,
		},
			RESPONSIBLE,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.plan)
		t.Run(testname, func(t *testing.T) {
			insuranceScore := NewRiskScore(tt.riskProfile.RiskScore)
			profile := NewInsuranceProfile(insuranceScore)
			autoRules := NewDisabilityRules()
			autoRules.evaluate(tt.riskProfile, profile)

			assert.Equal(t, tt.plan, profile.Disability.GetPlan())
		})
	}
}

func TestLifeRules(t *testing.T) {
	var tests = []struct {
		riskProfile RiskProfile
		plan        string
	}{
		{RiskProfile{
			Age:        41,
			Dependents: 1,
			House: func() *House {
				house := House{OwnershipStatus: OWNED}
				return &house
			}(),
			Income:        199999,
			MartialStatus: MARRIED,
			Vehicle: func() *Vehicle {
				vehicle := Vehicle{Year: 2010}
				return &vehicle
			}(),
			RiskScore: 1,
		},
			RESPONSIBLE,
		},
		{RiskProfile{
			Age:        40,
			Dependents: 0,
			House: func() *House {
				house := House{OwnershipStatus: OWNED}
				return &house
			}(),
			Income:        200001,
			MartialStatus: SINGLE,
			Vehicle: func() *Vehicle {
				vehicle := Vehicle{Year: 2010}
				return &vehicle
			}(),
			RiskScore: 0,
		},
			ECONOMIC,
		},
		{RiskProfile{
			Age:        61,
			Dependents: 0,
			House: func() *House {
				house := House{OwnershipStatus: OWNED}
				return &house
			}(),
			Income:        200001,
			MartialStatus: SINGLE,
			Vehicle: func() *Vehicle {
				vehicle := Vehicle{Year: 2010}
				return &vehicle
			}(),
			RiskScore: 0,
		},
			INELIGIBLE,
		},
		{RiskProfile{
			Age:        29,
			Dependents: 1,
			House: func() *House {
				house := House{OwnershipStatus: OWNED}
				return &house
			}(),
			Income:        150000,
			MartialStatus: MARRIED,
			Vehicle: func() *Vehicle {
				vehicle := Vehicle{Year: 2010}
				return &vehicle
			}(),
			RiskScore: 3,
		},
			RESPONSIBLE,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.plan)
		t.Run(testname, func(t *testing.T) {
			insuranceScore := NewRiskScore(tt.riskProfile.RiskScore)
			profile := NewInsuranceProfile(insuranceScore)
			autoRules := NewLifeRules()
			autoRules.evaluate(tt.riskProfile, profile)

			assert.Equal(t, tt.plan, profile.Life.GetPlan())
		})
	}
}

func TestHomeRules(t *testing.T) {
	var tests = []struct {
		riskProfile RiskProfile
		plan        string
	}{
		{RiskProfile{
			Age:           41,
			Dependents:    1,
			Income:        199999,
			MartialStatus: MARRIED,
			Vehicle: func() *Vehicle {
				vehicle := Vehicle{Year: 2010}
				return &vehicle
			}(),
			RiskScore: 1,
		},
			INELIGIBLE,
		},
		{RiskProfile{
			Age:        29,
			Dependents: 1,
			House: func() *House {
				house := House{OwnershipStatus: OWNED}
				return &house
			}(),
			Income:        199999,
			MartialStatus: MARRIED,
			Vehicle: func() *Vehicle {
				vehicle := Vehicle{Year: 2010}
				return &vehicle
			}(),
			RiskScore: 1,
		},
			ECONOMIC,
		},
		{RiskProfile{
			Age:        31,
			Dependents: 0,
			House: func() *House {
				house := House{OwnershipStatus: MORTGAGED}
				return &house
			}(),
			Income:        199999,
			MartialStatus: SINGLE,
			RiskScore:     3,
		},
			RESPONSIBLE,
		},
		{RiskProfile{
			Age:        41,
			Dependents: 0,
			House: func() *House {
				house := House{OwnershipStatus: MORTGAGED}
				return &house
			}(),
			Income:        200001,
			MartialStatus: MARRIED,
			RiskScore:     0,
		},
			ECONOMIC,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.plan)
		t.Run(testname, func(t *testing.T) {
			insuranceScore := NewRiskScore(tt.riskProfile.RiskScore)
			profile := NewInsuranceProfile(insuranceScore)
			autoRules := NewHomeRules()
			autoRules.evaluate(tt.riskProfile, profile)

			assert.Equal(t, tt.plan, profile.Home.GetPlan())
		})
	}
}
