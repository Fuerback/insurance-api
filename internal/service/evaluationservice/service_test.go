package evaluationservice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	service InsuranceEvaluation
)

func setupTest() {
	rules := []Rules{
		NewAutoRules(),
		NewHomeRules(),
		NewDisabilityRules(),
		NewLifeRules(),
	}

	service = NewService(rules)
}

func TestInsurancePlans(t *testing.T) {
	setupTest()

	var tests = []struct {
		userInformation                              UserInformation
		lifePlan, homePlan, autoPlan, disabilityPlan string
	}{
		{
			UserInformation{
				Age:        35,
				Dependents: 1,
				House: func() *House {
					house := House{OwnershipStatus: OWNED}
					return &house
				}(),
				Income:        100000,
				MartialStatus: SINGLE,
				RiskQuestions: []int{0, 0, 0},
				Vehicle: func() *Vehicle {
					vehicle := Vehicle{Year: 2020}
					return &vehicle
				}(),
			},
			ECONOMIC,
			ECONOMIC,
			ECONOMIC,
			ECONOMIC,
		},
		{
			UserInformation{
				Age:           45,
				Dependents:    0,
				Income:        0,
				MartialStatus: SINGLE,
				RiskQuestions: []int{0, 0, 0},
			},
			ECONOMIC,
			INELIGIBLE,
			INELIGIBLE,
			INELIGIBLE,
		},
		{
			UserInformation{
				Age:           45,
				Dependents:    0,
				Income:        0,
				MartialStatus: MARRIED,
				RiskQuestions: []int{1, 1, 1},
			},
			RESPONSIBLE,
			INELIGIBLE,
			INELIGIBLE,
			INELIGIBLE,
		},
		{
			UserInformation{
				Age:        35,
				Dependents: 1,
				House: func() *House {
					house := House{OwnershipStatus: MORTGAGED}
					return &house
				}(),
				Income:        200001,
				MartialStatus: MARRIED,
				RiskQuestions: []int{0, 0, 1},
				Vehicle: func() *Vehicle {
					vehicle := Vehicle{Year: 2013}
					return &vehicle
				}(),
			},
			REGULAR,
			ECONOMIC,
			ECONOMIC,
			ECONOMIC,
		},
		{
			UserInformation{
				Age:           65,
				Dependents:    0,
				Income:        200001,
				MartialStatus: SINGLE,
				RiskQuestions: []int{0, 1, 1},
				Vehicle: func() *Vehicle {
					vehicle := Vehicle{Year: 2022}
					return &vehicle
				}(),
			},
			INELIGIBLE,
			INELIGIBLE,
			REGULAR,
			INELIGIBLE,
		},
		{
			UserInformation{},
			ECONOMIC,
			INELIGIBLE,
			INELIGIBLE,
			INELIGIBLE,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s,%s,%s", tt.lifePlan, tt.homePlan, tt.autoPlan, tt.disabilityPlan)
		t.Run(testname, func(t *testing.T) {
			suggest := service.Evaluate(tt.userInformation)
			assert.Equal(t, tt.lifePlan, suggest.Life)
			assert.Equal(t, tt.homePlan, suggest.Home)
			assert.Equal(t, tt.autoPlan, suggest.Auto)
			assert.Equal(t, tt.disabilityPlan, suggest.Disability)
		})
	}
}
