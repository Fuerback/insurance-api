package evaluationservice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	rules []Rules
)

func loadAllRules() {
	rules = make([]Rules, 0)
	rules = append(rules, NewAutoRules())
	rules = append(rules, NewHomeRules())
	rules = append(rules, NewDisabilityRules())
	rules = append(rules, NewLifeRules())
}

func TestEvaluation(t *testing.T) {
	userInformation := UserInformation{
		Age:        10,
		Dependents: 1,
		House: func() *House {
			house := House{OwnershipStatus: "owned"}
			return &house
		}(),
		Income:        100000,
		MartialStatus: "single",
		RiskQuestions: []int8{1, 0, 0},
		Vehicle: func() *Vehicle {
			vehicle := Vehicle{Year: 2020}
			return &vehicle
		}(),
	}

	loadAllRules()
	s := NewService(rules)
	suggest := s.Evaluate(userInformation)

	assert.Equal(t, "economic", suggest.Life)
	assert.Equal(t, "economic", suggest.Auto)
	assert.Equal(t, "economic", suggest.Disability)
	assert.Equal(t, "economic", suggest.Home)
}
