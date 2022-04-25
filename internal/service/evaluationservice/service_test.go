package evaluationservice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"useorigin.com/insurance-api/internal/service/rules"
)

var (
	service InsuranceEvaluation
)

func setupTest() {
	service = NewService()
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
					house := House{OwnershipStatus: rules.OWNED}
					return &house
				}(),
				Income:        100000,
				MartialStatus: rules.SINGLE,
				RiskQuestions: []int{0, 0, 0},
				Vehicle: func() *Vehicle {
					vehicle := Vehicle{Year: 2020}
					return &vehicle
				}(),
			},
			rules.ECONOMIC,
			rules.ECONOMIC,
			rules.ECONOMIC,
			rules.ECONOMIC,
		},
		{
			UserInformation{
				Age:           45,
				Dependents:    0,
				Income:        0,
				MartialStatus: rules.SINGLE,
				RiskQuestions: []int{0, 0, 0},
			},
			rules.ECONOMIC,
			rules.INELIGIBLE,
			rules.INELIGIBLE,
			rules.INELIGIBLE,
		},
		{
			UserInformation{
				Age:           45,
				Dependents:    0,
				Income:        0,
				MartialStatus: rules.MARRIED,
				RiskQuestions: []int{1, 1, 1},
			},
			rules.RESPONSIBLE,
			rules.INELIGIBLE,
			rules.INELIGIBLE,
			rules.INELIGIBLE,
		},
		{
			UserInformation{
				Age:        35,
				Dependents: 1,
				House: func() *House {
					house := House{OwnershipStatus: rules.MORTGAGED}
					return &house
				}(),
				Income:        200001,
				MartialStatus: rules.MARRIED,
				RiskQuestions: []int{0, 0, 1},
				Vehicle: func() *Vehicle {
					vehicle := Vehicle{Year: 2013}
					return &vehicle
				}(),
			},
			rules.REGULAR,
			rules.ECONOMIC,
			rules.ECONOMIC,
			rules.ECONOMIC,
		},
		{
			UserInformation{
				Age:           65,
				Dependents:    0,
				Income:        200001,
				MartialStatus: rules.SINGLE,
				RiskQuestions: []int{0, 1, 1},
				Vehicle: func() *Vehicle {
					vehicle := Vehicle{Year: 2022}
					return &vehicle
				}(),
			},
			rules.INELIGIBLE,
			rules.INELIGIBLE,
			rules.REGULAR,
			rules.INELIGIBLE,
		},
		{
			UserInformation{},
			rules.ECONOMIC,
			rules.INELIGIBLE,
			rules.INELIGIBLE,
			rules.INELIGIBLE,
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
