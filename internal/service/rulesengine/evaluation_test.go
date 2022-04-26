package rulesengine

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsurancePlans(t *testing.T) {
	engine := NewEngine()
	var tests = []struct {
		riskProfile                                  RiskProfile
		lifePlan, homePlan, autoPlan, disabilityPlan string
	}{
		{
			RiskProfile{
				Age:           20,
				Dependents:    0,
				Income:        150000,
				MartialStatus: SINGLE,
				RiskScore:     3,
			},
			REGULAR,
			INELIGIBLE,
			INELIGIBLE,
			REGULAR,
		},
		{
			RiskProfile{
				Age:        20,
				Dependents: 0,
				House: func() *House {
					house := House{OwnershipStatus: MORTGAGED}
					return &house
				}(),
				Income:        150000,
				MartialStatus: SINGLE,
				Vehicle: func() *Vehicle {
					vehicle := Vehicle{Year: 2020}
					return &vehicle
				}(),
				RiskScore: 1,
			},
			ECONOMIC,
			ECONOMIC,
			ECONOMIC,
			ECONOMIC,
		},
		{
			RiskProfile{
				Age:        20,
				Dependents: 0,
				House: func() *House {
					house := House{OwnershipStatus: MORTGAGED}
					return &house
				}(),
				Income:        150000,
				MartialStatus: SINGLE,
				Vehicle: func() *Vehicle {
					vehicle := Vehicle{Year: 2020}
					return &vehicle
				}(),
				RiskScore: 0,
			},
			ECONOMIC,
			ECONOMIC,
			ECONOMIC,
			ECONOMIC,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s,%s,%s", tt.lifePlan, tt.homePlan, tt.autoPlan, tt.disabilityPlan)
		t.Run(testname, func(t *testing.T) {
			suggest := engine.EvaluateRules(tt.riskProfile)
			assert.Equal(t, tt.lifePlan, suggest.Life.GetPlan())
			assert.Equal(t, tt.homePlan, suggest.Home.GetPlan())
			assert.Equal(t, tt.autoPlan, suggest.Auto.GetPlan())
			assert.Equal(t, tt.disabilityPlan, suggest.Disability.GetPlan())
		})
	}
}
