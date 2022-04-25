package evaluationservice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	rules   []Rules
	service InsuranceEvaluation
)

func setupTest() {
	rules = make([]Rules, 0)
	rules = append(rules, NewAutoRules())
	rules = append(rules, NewHomeRules())
	rules = append(rules, NewDisabilityRules())
	rules = append(rules, NewLifeRules())

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
					house := House{OwnershipStatus: "owned"}
					return &house
				}(),
				Income:        100000,
				MartialStatus: "single",
				RiskQuestions: []int{0, 0, 0},
				Vehicle: func() *Vehicle {
					vehicle := Vehicle{Year: 2020}
					return &vehicle
				}(),
			},
			"economic",
			"economic",
			"economic",
			"economic",
		},
		{
			UserInformation{
				Age:           45,
				Dependents:    0,
				Income:        0,
				MartialStatus: "single",
				RiskQuestions: []int{0, 0, 0},
			},
			"economic",
			"ineligible",
			"ineligible",
			"ineligible",
		},
		{
			UserInformation{
				Age:           45,
				Dependents:    0,
				Income:        0,
				MartialStatus: "married",
				RiskQuestions: []int{1, 1, 1},
			},
			"responsible",
			"ineligible",
			"ineligible",
			"ineligible",
		},
		{
			UserInformation{
				Age:        35,
				Dependents: 1,
				House: func() *House {
					house := House{OwnershipStatus: "mortgaged"}
					return &house
				}(),
				Income:        200001,
				MartialStatus: "married",
				RiskQuestions: []int{0, 0, 1},
				Vehicle: func() *Vehicle {
					vehicle := Vehicle{Year: 2013}
					return &vehicle
				}(),
			},
			"regular",
			"economic",
			"economic",
			"economic",
		},
		{
			UserInformation{
				Age:           20,
				Dependents:    0,
				Income:        200001,
				MartialStatus: "single",
				RiskQuestions: []int{0, 1, 1},
				Vehicle: func() *Vehicle {
					vehicle := Vehicle{Year: 2022}
					return &vehicle
				}(),
			},
			"economic",
			"ineligible",
			"economic",
			"economic",
		},
		{
			UserInformation{},
			"economic",
			"ineligible",
			"ineligible",
			"ineligible",
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
