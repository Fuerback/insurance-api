package rulesengine

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsurancePlans(t *testing.T) {
	var tests = []struct {
		evaluation                                   *evaluation
		riskProfile                                  RiskProfile
		lifePlan, homePlan, autoPlan, disabilityPlan string
	}{
		{
			NewEvaluation([]Rule{NewDisabilityRules(), NewLifeRules()}),
			RiskProfile{
				Age:           20,
				Dependents:    0,
				Income:        150000,
				MartialStatus: SINGLE,
				RiskScore:     3,
			},
			REGULAR,
			RESPONSIBLE,
			RESPONSIBLE,
			REGULAR,
		},
		{
			NewEvaluation([]Rule{NewDisabilityRules(), NewLifeRules(), NewAutoRules(), NewHomeRules()}),
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
			NewEvaluation([]Rule{}),
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
			suggest := tt.evaluation.EvaluateRules(tt.riskProfile)
			assert.Equal(t, tt.lifePlan, suggest.Life.GetPlan())
			assert.Equal(t, tt.homePlan, suggest.Home.GetPlan())
			assert.Equal(t, tt.autoPlan, suggest.Auto.GetPlan())
			assert.Equal(t, tt.disabilityPlan, suggest.Disability.GetPlan())
		})
	}
}
