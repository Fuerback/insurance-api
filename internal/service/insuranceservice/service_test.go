package insuranceservice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"useorigin.com/insurance-api/internal/service/rulesengine"
)

var (
	service Insurance
)

func setupTest() {
	service = NewService()
}

func TestInsurancePlans(t *testing.T) {
	setupTest()

	var tests = []struct {
		userInformation                              RiskProfile
		lifePlan, homePlan, autoPlan, disabilityPlan string
	}{
		{
			RiskProfile{
				Age:        35,
				Dependents: 1,
				House: func() *House {
					house := House{OwnershipStatus: rulesengine.OWNED}
					return &house
				}(),
				Income:        100000,
				MartialStatus: rulesengine.SINGLE,
				RiskQuestions: []int{0, 0, 0},
				Vehicle: func() *Vehicle {
					vehicle := Vehicle{Year: 2020}
					return &vehicle
				}(),
			},
			rulesengine.ECONOMIC,
			rulesengine.ECONOMIC,
			rulesengine.ECONOMIC,
			rulesengine.ECONOMIC,
		},
		{
			RiskProfile{
				Age:           45,
				Dependents:    0,
				Income:        0,
				MartialStatus: rulesengine.SINGLE,
				RiskQuestions: []int{0, 0, 0},
			},
			rulesengine.ECONOMIC,
			rulesengine.INELIGIBLE,
			rulesengine.INELIGIBLE,
			rulesengine.INELIGIBLE,
		},
		{
			RiskProfile{
				Age:           45,
				Dependents:    0,
				Income:        0,
				MartialStatus: rulesengine.MARRIED,
				RiskQuestions: []int{1, 1, 1},
			},
			rulesengine.RESPONSIBLE,
			rulesengine.INELIGIBLE,
			rulesengine.INELIGIBLE,
			rulesengine.INELIGIBLE,
		},
		{
			RiskProfile{
				Age:        35,
				Dependents: 1,
				House: func() *House {
					house := House{OwnershipStatus: rulesengine.MORTGAGED}
					return &house
				}(),
				Income:        200001,
				MartialStatus: rulesengine.MARRIED,
				RiskQuestions: []int{0, 0, 1},
				Vehicle: func() *Vehicle {
					vehicle := Vehicle{Year: 2013}
					return &vehicle
				}(),
			},
			rulesengine.REGULAR,
			rulesengine.ECONOMIC,
			rulesengine.ECONOMIC,
			rulesengine.ECONOMIC,
		},
		{
			RiskProfile{
				Age:           65,
				Dependents:    0,
				Income:        200001,
				MartialStatus: rulesengine.SINGLE,
				RiskQuestions: []int{0, 1, 1},
				Vehicle: func() *Vehicle {
					vehicle := Vehicle{Year: 2022}
					return &vehicle
				}(),
			},
			rulesengine.INELIGIBLE,
			rulesengine.INELIGIBLE,
			rulesengine.REGULAR,
			rulesengine.INELIGIBLE,
		},
		{
			RiskProfile{},
			rulesengine.ECONOMIC,
			rulesengine.INELIGIBLE,
			rulesengine.INELIGIBLE,
			rulesengine.INELIGIBLE,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s,%s,%s", tt.lifePlan, tt.homePlan, tt.autoPlan, tt.disabilityPlan)
		t.Run(testname, func(t *testing.T) {
			suggest := service.EvaluateUserProfile(tt.userInformation)
			assert.Equal(t, tt.lifePlan, suggest.Life)
			assert.Equal(t, tt.homePlan, suggest.Home)
			assert.Equal(t, tt.autoPlan, suggest.Auto)
			assert.Equal(t, tt.disabilityPlan, suggest.Disability)
		})
	}
}
